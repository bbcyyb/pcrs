package middlewares

import (
	"github.com/bbcyyb/pcrs/infra/log"
	"github.com/spf13/viper"
)

type config struct {
	JWTEnable  bool
	LogEnable  bool
	AuthEnable bool
}

var C *config

func initConfig() {
	log.Info("Initalize config info for middlewares package")

	C = &config{
		JWTEnable:  viper.GetBool("middlewares.jwt.enable"),
		LogEnable:  viper.GetBool("middlewares.Log.enable"),
		AuthEnable: viper.GetBool("middlewares.authorization.enable"),
	}

	log.Debug("middlewares Config is ", C)
}

func Setup() {
	initConfig()
}
