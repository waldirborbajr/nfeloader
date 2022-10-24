package util

import (
	"os"

	"github.com/rs/zerolog/log"
)

var hasErr = false
var errMsg = ""

func Healthcheck(path string) string {
	log.Info().Msg("HealthCheking Folders [XMLS, LOGS & PROCESSED]")

	folders := []string{"xmls", "logs", "processed", "xmlerror"}

	for _, folder := range folders {

		folderPath := path + "/" + folder

		_, err := os.Stat(folderPath)

		if os.IsNotExist(err) {
			err := os.Mkdir(folderPath, 0o755)

			if err != nil {
				hasErr = true
				errMsg += err.Error() + " "
				continue

			}
		}
	}

	if hasErr {
		return errMsg
	}

	return ""
}
