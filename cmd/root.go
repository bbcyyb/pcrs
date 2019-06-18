package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "pcrs",
	Short: "PowerCalculator Restful Service",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {

	},
}
