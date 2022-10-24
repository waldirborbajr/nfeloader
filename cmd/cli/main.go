package main

import (
	"os"
	"sync"

	"github.com/rs/zerolog/log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/waldirborbajr/nfeloader/internal/cronjob"
	"github.com/waldirborbajr/nfeloader/internal/util"
	"github.com/waldirborbajr/nfeloader/pkg/config"
	"github.com/waldirborbajr/nfeloader/pkg/customlog"
	"github.com/waldirborbajr/nfeloader/pkg/service"
	"github.com/waldirborbajr/nfeloader/pkg/xml"
)

func init() {
	log.Info().Msg(".\n")
	log.Info().Msg("Starting NFeLoader " + config.Verzion + "\n")

	config.AppPath, config.Err = os.Getwd()

	util.Healthcheck(config.AppPath)

	if config.Err != nil {
		customlog.HandleError("Path", config.Err)
	}

	config.Cfg = config.BuildConfig()

	config.DBcon = dsn()
}

func dsn() string {
	return config.MysqlUrl(config.Cfg)
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

	err = xml.MoveXML(config.AppPath, file)

	if err != nil {
		customlog.HandleError("Moving XML", err)
	}
}

func main() {
	cronjob.RunCronJob()
	// mainNormal()
}
