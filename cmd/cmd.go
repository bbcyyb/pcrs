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

	if configFile != "" {
		viper.SetConfigFile(configFile)
		log.Info("Load config from ", configFile)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		viper.AddConfigPath("/etc/pcrs")
		viper.AddConfigPath("$HOME/.pcrs")
		log.Info("Load config from default path")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Error(err)
	}

	log.Info("Initialize config info for cmd package")

	port := viper.GetInt("port")
	log.Debug("Config port is ", port)
	C = &config{
		Port: port,
	}

	log.Debug("port is ", port)

	if C.Port == 0 {
		C.Port = 8080
	}
}

func init() {
	InitRoot()
}
