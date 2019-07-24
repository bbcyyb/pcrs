package cmd

import (
	"github.com/bbcyyb/pcrs/pkg/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/sys/unix"

	"github.com/bbcyyb/pcrs/conf"
	"github.com/bbcyyb/pcrs/middlewares"
	"github.com/bbcyyb/pcrs/pkg"
	"github.com/bbcyyb/pcrs/pkg/log"
)

var rootCmd = &cobra.Command{
	Use:   "pcrs",
	Short: "PowerCalculator Restful Service",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
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

	conf.Setup(configFile)
	pkg.Setup()
	middlewares.Setup()
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
