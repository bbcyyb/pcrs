package cmd

import (
	"github.com/bbcyyb/pcrs/infra/log"
	"github.com/golang-migrate/migrate/v4/internal/cli"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate database",
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("migrate running.")
		cli.Main("v1.0.0")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
