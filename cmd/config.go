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
	log.Info("Initialize config info for cmd package")

	port := viper.GetInt("port")

	C = &config{
		Port: port,
	}

	log.Debug("port is ", port)

	if C.Port == 0 {
		C.Port = 8080
	}
}
