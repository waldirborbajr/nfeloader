package main

import (
	"os"

	"github.com/rs/zerolog/log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/waldirborbajr/nfeloader/internal/config"
	"github.com/waldirborbajr/nfeloader/internal/cronjob"
	"github.com/waldirborbajr/nfeloader/internal/customlog"
	"github.com/waldirborbajr/nfeloader/internal/util"
	"github.com/waldirborbajr/nfeloader/internal/version"
)

func main() {
	log.Info().Msg("Starting NFeLoader " + version.AppVersion())

	config.AppPath, config.Err = os.Getwd()

	util.Healthcheck(config.AppPath)

	if config.Err != nil {
		customlog.HandleError("Path", config.Err)
	}

	config.Cfg = config.BuildConfig()

	config.DBcon = util.Dsn()

	cronjob.RunCronJob()
}
