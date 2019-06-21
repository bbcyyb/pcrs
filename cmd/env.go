package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Print pcrs version and environment info",
	Long:  `Print pcrs version and environment info. This is useful in pcrs bug reports.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GOOS=", runtime.GOOS)
		fmt.Println("GOARCH=", runtime.GOARCH)
		fmt.Println("GOVERSION=", runtime.Version())
	},
}

func init() {
	rootCmd.AddCommand(envCmd)
}
