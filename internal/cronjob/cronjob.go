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
	"github.com/waldirborbajr/nfeloader/pkg/config"
	"github.com/waldirborbajr/nfeloader/pkg/customlog"
	"github.com/waldirborbajr/nfeloader/pkg/mail"
	"github.com/waldirborbajr/nfeloader/pkg/repository"
	"github.com/waldirborbajr/nfeloader/pkg/service"
	"github.com/waldirborbajr/nfeloader/pkg/xml"
)

func RunNoCronJob() {
	// telegramAPI := telegram.NewAPI(cfg.TelegramChatID, cfg.TelegramBotToken)

	// var xmlFiles []string
	path := config.AppPath + "/xmls/"

	// file, err := os.OpenFile(
	// 	appPath+"/logs/nfeloader.log",
	// 	os.O_CREATE|os.O_APPEND|os.O_WRONLY,
	// 	0o644,
	// )
	// if err != nil {
	// 	customlog.HandleError("Creating log file", err)
	// 	os.Exit(-1)
	// }
	// log.SetOutput(file)

	log.Info().Msg("======================================================")
	log.Info().Msgf("Server v%s pid=%d started with processes: %d",
		config.Verzion, os.Getpid(), runtime.GOMAXPROCS(runtime.NumCPU()))

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
		customlog.HandleError("Listing XML", mErr)
	}

	if len(xmlFiles) != 0 {
		log.Info().Msgf("Found %d XML(s) file(s)", len(xmlFiles))
		db, dbErr := sql.Open("mysql", dbcon)
		if dbErr != nil {
			customlog.HandleError("Opening database connection", dbErr)
		} else {

			db.SetMaxOpenConns(10)
			db.SetMaxIdleConns(10)
			db.SetConnMaxLifetime(time.Second * 60)

			repositoryMysql := repository.NewNFeProcMysql(db)
			service := service.NewNFeProcService(repositoryMysql)

			wg := &sync.WaitGroup{}
			for _, file := range xmlFiles {
				wg.Add(1)
				go worker(wg, path, file, service)

			}
			wg.Wait()
		}
	}

	log.Info().Msgf("Finished %s", time.Now())
	log.Info().Msgf("Elapsed time  %s", time.Since(start))
}

func RunCronJob() {
	log.Info().Msg("\n🚀\n")

	tmz, _ := time.LoadLocation("America/Sao_Paulo")
	// c := cron.New(cron.WithChain(cron.SkipIfStillRunning(logger)))

	cr := cron.New(cron.WithLocation(tmz))

	log.Info().Msgf("Cronjob: %s", cfg.Schedule)

	sched := fmt.Sprintf("@every %s", cfg.Schedule)

	cr.AddFunc(sched, controledJob)

	cr.Start()

	for {
		time.Sleep(time.Second)
	}
}
