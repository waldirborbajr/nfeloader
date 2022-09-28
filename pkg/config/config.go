package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type NFeConfig struct {
	MailServer       string
	MailUsr          string
	MailPwd          string
	DatabaseHost     string
	DatabaseUsr      string
	DatabasePwd      string
	DatabaseDbName   string
	TelegramChatID   string
	TelegramBotToken string
	Schedule         string
}

func BuildConfig() *NFeConfig {
	cfg := &NFeConfig{}

	isContainer, _ := strconv.ParseBool(os.Getenv("CONTAINER"))

	log.Info().Msgf("isContainer: %v", isContainer)

	if !isContainer {
		err := godotenv.Load()
		if err != nil {
			log.Info().Msgf("ERROR: loading .env file: %v", err)
		}
	}

	cfg.MailServer = os.Getenv("MAIL_SERVER")
	cfg.MailUsr = os.Getenv("MAIL_USR")
	cfg.MailPwd = os.Getenv("MAIL_PWD")
	cfg.DatabaseHost = os.Getenv("DATABASE_HOST")
	cfg.DatabaseUsr = os.Getenv("DATABASE_USR")
	cfg.DatabasePwd = os.Getenv("DATABASE_PWD")
	cfg.DatabaseDbName = os.Getenv("DATABASE_NAME")
	cfg.TelegramChatID = os.Getenv("TELEGRAM_CHAT_ID")
	cfg.TelegramBotToken = os.Getenv("TELEGRAM_BOT_TOKEN")
	cfg.Schedule = os.Getenv("TIME_SCHEDULE")

	return cfg
}
