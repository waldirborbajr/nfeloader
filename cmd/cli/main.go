package main

import (
	"database/sql"
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/rs/zerolog/log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/robfig/cron/v3"
	"github.com/waldirborbajr/nfeloader/pkg/config"
	"github.com/waldirborbajr/nfeloader/pkg/customlog"
	"github.com/waldirborbajr/nfeloader/pkg/mail"
	"github.com/waldirborbajr/nfeloader/pkg/repository"
	"github.com/waldirborbajr/nfeloader/pkg/service"
	"github.com/waldirborbajr/nfeloader/pkg/xml"
)

var (
	dbcon   string
	err     error
	cfg     *config.NFeConfig
	appPath string
)

func init() {
	log.Info().Msg(".\n")
	log.Info().Msg("Starting NFeLoader " + config.Verzion + "\n")

	appPath, err = os.Getwd()

	healthcheck(appPath)

	if err != nil {
		customlog.HandleError("Path", err)
	}

	cfg = config.BuildConfig()

	dbcon = dsn()
}

func healthcheck(path string) {
	// log.Println("HealthCheking")
	log.Info().Msg("HealthCheking")

	folders := []string{"xmls", "logs", "processed"}

	for _, content := range folders {

		folderPath := path + "/" + content

		_, err := os.Stat(folderPath)

		if os.IsNotExist(err) {
			os.Mkdir(folderPath, 0o755)
		}
	}
}

func dsn() string {
	return config.MysqlUrl(cfg)
}

func worker(wg *sync.WaitGroup, path string, file string, service *service.NFeProcService) {
	defer wg.Done()

	nfeProc, err := xml.ReadXML(path, file)
	if err != nil {
		customlog.HandleError("Reading XML", err)
	}

	// Call service to save
	if err = service.SaveNFe(nfeProc); err != nil {
		customlog.HandleError("Saving NFe", err)
	}

	err = xml.MoveXML(appPath, file)

	if err != nil {
		customlog.HandleError("Moving XML", err)
	}
}

func controledJob() {
	// telegramAPI := telegram.NewAPI(cfg.TelegramChatID, cfg.TelegramBotToken)

	// var xmlFiles []string
	path := appPath + "/xmls/"

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
		version, os.Getpid(), runtime.GOMAXPROCS(runtime.NumCPU()))

	start := time.Now()
	log.Info().Msgf("Starting NF-e Loader: %s", time.Now())

	log.Info().Msgf("Searching for new e-mails")

	mErr := mail.NewMessage(path, cfg)

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

func mainCron() {
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

func mainNormal() {
	controledJob()
}

func main() {
	mainCron()
	// mainNormal()
}
