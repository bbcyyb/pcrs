package pkg

import (
	"github.com/bbcyyb/pcrs/pkg/log"
	"github.com/spf13/viper"
)

type Config struct {
	AuthPolicyFile string
	AuthModelFile  string
}

var C *Config

func initConfig() {
	log.Info("Initalize config info for pkg package")

	C = &Config{
		AuthPolicyFile: viper.GetString("middlewares.authorization.policy"),
		AuthModelFile:  viper.GetString("middlewares.authorization.model"),
	}

	log.Debug("pkg Config is ", C)
}

func Setup() {
	initConfig()
}
