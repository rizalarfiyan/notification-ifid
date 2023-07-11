package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/rizalarfiyan/notification-ifid/utils"
)

type Config struct {
	Port           int
	TelegramToken  string
	TelegramChatId string
}

var conf *Config

func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(".env is not loaded properly")
	}

	conf = new(Config)
	conf.Port = utils.GetEnvAsInt("PORT", 8080)
	conf.TelegramToken = utils.GetEnv("TELEGRAM_TOKEN", "")
	conf.TelegramChatId = utils.GetEnv("TELEGRAM_CHAT_ID", "")
}

func Get() *Config {
	return conf
}
