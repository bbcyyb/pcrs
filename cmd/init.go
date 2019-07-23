package cmd

import (
	"github.com/bbcyyb/pcrs/pkg/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/sys/unix"

	"github.com/bbcyyb/pcrs/middlewares"
	"github.com/bbcyyb/pcrs/pkg"
	"github.com/bbcyyb/pcrs/pkg/log"
)

type config struct {
	Port int
}

var C *config

var rootCmd = &cobra.Command{
	Use:   "pcrs",
	Short: "PowerCalculator Restful Service",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
}

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		viper.AddConfigPath("/etc/todos")
		viper.AddConfigPath("$HOME/.pcrs")

		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			log.Error(err)
		}
	}

	log.Info("Initialize config info for cmd package")

	port := viper.GetInt("database.port")
	log.Debug("Config Node port is ", port)
	C = &config{
		Port: port,
	}

	if C.Port == 0 {
		C.Port = 8080
	}
}

func initLog() {
	if !terminal.IsTerminal(unix.Stdout) {
		log.SetFormatter(log.JSON)
	} else {
		log.SetFormatter(log.TEXT)
	}

	if verbose, _ := rootCmd.Flags().GetBool("verbose"); verbose {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	log.Info("Log settings is ready")
}

func Setup() {
	initLog()
	initConfig()

	middlewares.Setup()
	pkg.Setup()
	logger.Setup()

	//	Execute()
}

func Execute() (err error) {
	err = rootCmd.Execute()
	return
}

var configFile string

func init() {
	cobra.OnInitialize(Setup)
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "make output more verbose")
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is config.yaml)")
}
