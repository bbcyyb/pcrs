package infra

import (
	"github.com/bbcyyb/pcrs/infra/log"
	"github.com/spf13/viper"
)

type config struct {
	AuthPolicyFile string
	AuthModelFile  string
}

var C *config

func initConfig() {
	log.Info("Initalize config info for infra package")

	C = &config{
		AuthPolicyFile: viper.GetString("middlewares.authorization.policy"),
		AuthModelFile:  viper.GetString("middlewares.authorization.model"),
	}

	log.Debug("infra Config is ", C)
}

func Setup() {
	initConfig()
}
