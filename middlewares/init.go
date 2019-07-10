package middlewares

import (
	"github.com/bbcyyb/pcrs/infra/log"
	"github.com/spf13/viper"
)

type config struct {
	AuthTEnable bool
	AuthREnable bool
	LogEnable   bool
}

var C *config

func initConfig() {
	log.Info("Initalize config info for middlewares package")

	C = &config{
		AuthTEnable: viper.GetBool("middlewares.authentication.enable"),
		AuthREnable: viper.GetBool("middlewares.authorization.enable"),
		LogEnable:   viper.GetBool("middlewares.Log.enable"),
	}

	log.Debug("middlewares Config is ", C)
}

func Setup() {
	initConfig()
}
