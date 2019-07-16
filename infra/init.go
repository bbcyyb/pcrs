package infra

import (
	"github.com/bbcyyb/pcrs/infra/log"
	"github.com/spf13/viper"
)

type Config struct {
	AuthPolicyFile string
	AuthModelFile  string
}

var C *Config

func initConfig() {
	log.Info("Initalize config info for infra package")

	C = &Config{
		AuthPolicyFile: viper.GetString("middlewares.authorization.policy"),
		AuthModelFile:  viper.GetString("middlewares.authorization.model"),
	}

	log.Debug("infra Config is ", C)
}

func Setup() {
	initConfig()
}
