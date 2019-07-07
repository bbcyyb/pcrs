package middlewares

import (
	"github.com/bbcyyb/pcrs/infra/log"
	"github.com/spf13/viper"
)

type config struct {
	JWT bool
}

var C *config

func initConfig() {
	log.Info("Initalize config info for middlewares package")
	jwt := viper.GetBool("middlewares.jwt")

	C = &config{
		JWT: jwt,
	}

	log.Debug("Config Node middlewares.jwt is ", jwt)
}

func Setup() {
	initConfig()
}
