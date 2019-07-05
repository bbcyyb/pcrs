package cmd

import (
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/sys/unix"

	"github.com/bbcyyb/pcrs/infra/log"
)

var rootCmd = &cobra.Command{
	Use:   "pcrs",
	Short: "PowerCalculator Restful Service",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() (err error) {
	err = rootCmd.Execute()
	return
}

func initLogSettings() {
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

var configFile string

func InitRoot() {
	cobra.OnInitialize(func() {
		initLogSettings()
		InitConfig()
	})
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "make output more verbose")
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is config.yaml)")
}
