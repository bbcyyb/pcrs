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
	log.SetFormatter(log.JSON)
	log.SetLevel(log.DebugLevel)
	log.Errorln("Errorln")
	log.Warnln("Warnln")
	log.Infoln("Infoln")
	log.Debugln("Debugln")
	log.Traceln("Traceln")
	fmt.Println("hello world, ", configFile)
	/*
		slice, err := log.Refresh()
		if err != nil {
			fmt.Println(err)
		}

		for _, val := range slice {
			fmt.Println(val)
		}
	*/
}
