package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/sys/unix"

	"github.com/bbcyyb/pcrs/infra/log"
)

var rootCmd = &cobra.Command{
	Use:   "pcrs",
	Short: "PowerCalculator Restful Service",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if !terminal.IsTerminal(unix.Stdout) {
			log.SetFormatter(log.JSON)
		} else {
			log.SetFormatter(log.TEXT)
		}

		if verbose, _ := cmd.Flags().GetBool("verbose"); verbose {
			log.SetLevel(log.DebugLevel)
		} else {
			log.SetLevel(log.InfoLevel)
		}
	},
}

func Execute() (err error) {
	err = rootCmd.Execute()
	return
}

var configFile string

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		viper.AddConfigPath("/etc/pcrs")
		viper.AddConfigPath("$HOME/.pcrs")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Error(err)
	}

	InitConfig()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "make output more verbose")
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is config.yaml)")
}
