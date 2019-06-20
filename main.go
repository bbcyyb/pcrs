package main

import (
	"fmt"

	"github.com/spf13/viper"

	log "github.com/bbcyyb/pcrs/infra/log"
)

var configFile string

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	}
}

func main() {
	configFile = "config.yaml"
	initConfig()
	log.InitLog()
	fmt.Println("hello world, ", configFile)
}
