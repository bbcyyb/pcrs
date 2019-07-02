package cmd

import (
	"github.com/bbcyyb/pcrs/infra/log"
	"github.com/spf13/viper"
)

type config struct {
	Port int
}

var C *config

func InitConfig() {
	log.Debug("Initialize config info for cmd package")

	C = &config{
		Port: viper.GetInt("port"),
	}

	log.Debug("port is ", viper.GetInt("port"))

	if C.Port == 0 {
		C.Port = 8080
	}
}
