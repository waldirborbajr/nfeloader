package main

import (
	"os"

	"github.com/rs/zerolog/log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/waldirborbajr/nfeloader/internal/cronjob"
	"github.com/waldirborbajr/nfeloader/internal/util"
	"github.com/waldirborbajr/nfeloader/pkg/config"
	"github.com/waldirborbajr/nfeloader/pkg/customlog"
)

func init() {
	log.Info().Msg("Starting NFeLoader " + config.Verzion + "\n")

	config.AppPath, config.Err = os.Getwd()

	util.Healthcheck(config.AppPath)

	if config.Err != nil {
		customlog.HandleError("Path", config.Err)
	}

	config.Cfg = config.BuildConfig()

	config.DBcon = util.Dsn()
}

func main() {
	cronjob.RunCronJob()
}
