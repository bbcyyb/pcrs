package main

import (
	"fmt"

	"github.com/spf13/viper"
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
	fmt.Println("hello world, ", configFile)
}
