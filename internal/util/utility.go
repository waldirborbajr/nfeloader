package util

import (
	"os"

	"github.com/rs/zerolog/log"
)

var hasErr = false
var errMsg = ""

func Healthcheck(path string) string {
	log.Info().Msg("HealthCheking")

	folders := []string{"xmls", "logs", "processed"}

	for _, content := range folders {

		folderPath := path + "/" + content

		_, err := os.Stat(folderPath)

		if os.IsNotExist(err) {
			err := os.Mkdir(folderPath, 0o755)

			if err != nil {
				hasErr = true
				errMsg = errMsg + " "
				continue

			}
		}
	}

	if hasErr {
		return errMsg
	}

	return ""
}
