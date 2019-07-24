package conf

import (
	"github.com/spf13/viper"
)

type Config struct {
	DB         Database   `mapstructure: "db"`
	Middleware Middleware `mapstructure: "middleware"`
	Pkg        Pkg        `mapstructure: "pkg"`
}

type Database struct {
	Server   string `mapstructure:"server"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	Source   string `mapstructure:"source"`
}

type Middleware struct {
	Authentication AuthT        `mapstructure: "authentication"`
	Authorization  AuthZ        `mapstructure: "authorization"`
	ErrorHandler   ErrorHandler `mapstructure: "errorhandler"`
}

type AuthT struct {
	Enable bool `mapstructure: "enable"`
}

type AuthZ struct {
	Enable bool `mapstructure: "enable"`
}

type ErrorHandler struct {
	Enable bool `mapstructure: "enable"`
}

type Pkg struct {
	Authorizer Authorizer `mapstructure: "authorizer"`
	Log        Log        `mapstructure: log`
}

type Authorizer struct {
	Policy string `mapstructure: "policy"`
	Model  string `mapstructure: "model"`
}

type Log struct {
	Level     string `mapstructure: "level"`
	Formatter string `mapstructure: "formatter"`
	LineNum   bool   `mapstructure: "linenum"`
	FileName  string `mapstructure: "filename"`
}

func initConfig(configFile string, c interface{}) {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		viper.AddConfigPath("./conf")

	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(c); err != nil {
		panic(err)
	}

}

var C Config

func Setup(configFile string) {
	initConfig(configFile, &C)
}
