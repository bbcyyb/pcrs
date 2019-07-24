package cmd

import (
	"github.com/bbcyyb/pcrs/conf"
	"github.com/bbcyyb/pcrs/middlewares"
	"github.com/bbcyyb/pcrs/pkg"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pcrs",
	Short: "PowerCalculator Restful Service",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
}

func Setup() {
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
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is config.yaml)")
}
