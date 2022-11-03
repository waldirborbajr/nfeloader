package cronjob

import (
	"database/sql"
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog/log"
	"github.com/waldirborbajr/nfeloader/internal/config"
	"github.com/waldirborbajr/nfeloader/internal/customlog"
	"github.com/waldirborbajr/nfeloader/internal/mail"
	"github.com/waldirborbajr/nfeloader/internal/repository"
	"github.com/waldirborbajr/nfeloader/internal/service"
	"github.com/waldirborbajr/nfeloader/internal/version"
	"github.com/waldirborbajr/nfeloader/internal/xml"
)

var wg sync.WaitGroup

func mainJob() {
	// var xmlFiles []string
	path := config.AppPath + "/xmls/"

	log.Info().Msg("======================================================")
	log.Info().Msgf("Server %s pid=%d started with processes: %d",
		version.AppVersion(), os.Getpid(), runtime.GOMAXPROCS(runtime.NumCPU()))

	start := time.Now()
	log.Info().Msgf("Starting NF-e Loader: %s", time.Now())

	log.Info().Msgf("Searching for new e-mails")

	mErr := mail.NewMessage(path, config.Cfg)

	if mErr != nil {
		customlog.HandleError("Verifying e-mail", mErr)
	}

	log.Info().Msg("Searching for XMLs files...")

	xmlFiles, err := xml.ListXML(path)
	if err != nil {
		customlog.HandleError("Listing XML %v", err)
	}

	if len(xmlFiles) != 0 {
		log.Info().Msgf("Found %d XML(s) file(s)", len(xmlFiles))
		db, dbErr := sql.Open("mysql", config.DBcon)
		if dbErr != nil {
			customlog.HandleError("Opening database connection", dbErr)
		} else {

			db.SetMaxOpenConns(10)
			db.SetMaxIdleConns(10)
			db.SetConnMaxLifetime(time.Second * 60)

			repositoryMysql := repository.NewNFeProcMysql(db)
			service := service.NewNFeProcService(repositoryMysql)

			semaphore := make(chan struct{}, 20)

			for _, file := range xmlFiles {
				wg.Add(1)
				semaphore <- struct{}{}
				go processXMLFile(path, file, service, semaphore)

			}

			wg.Wait()
		}
	}

	log.Info().Msgf("Finished %s", time.Now())
	log.Info().Msgf("Elapsed time  %s", time.Since(start))
}

func processXMLFile(path string, file string, service *service.NFeProcService, semaphore <-chan struct{}) {
	defer wg.Done()

	nfeProc, err := xml.ReadXML(path, file)
	if err != nil {
		<-semaphore
		customlog.HandleError("Error processing [ "+file+"] :", err)

		err = xml.MoveXML(config.AppPath, file, true)
		if err != nil {
			customlog.HandleError("Moving XML", err)
		}
	}

	// Call service to save
	if err = service.SaveNFe(nfeProc); err != nil {
		<-semaphore
		customlog.HandleError("Saving NFe", err)
	} else {
		err = xml.MoveXML(config.AppPath, file, false)

		if err != nil {
			customlog.HandleError("Moving XML", err)
		}
	}
	<-semaphore

}

func RunCronJob() {
	log.Info().Msg("🚀 CronJob")

	tmz, _ := time.LoadLocation("America/Sao_Paulo")

	cr := cron.New(cron.WithLocation(tmz))

	log.Info().Msgf("Cronjob: %s", config.Cfg.Schedule)

	sched := fmt.Sprintf("@every %s", config.Cfg.Schedule)

	_, err := cr.AddFunc(sched, mainJob)
	if err != nil {
		log.Info().Msgf("Error -> Cronjob: AddFunc")
	}

	cr.Start()

	for {
		time.Sleep(time.Second)
	}
}
