package cmd

import (
	"github.com/bbcyyb/pcrs/controllers"
	"github.com/bbcyyb/pcrs/middlewares"
	"github.com/bbcyyb/pcrs/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serves the api",
	Run: func(cmd *cobra.Command, args []string) {
		if verbose, _ := cmd.Flags().GetBool("verbose"); verbose {
			gin.SetMode(gin.DebugMode)
		} else {
			gin.SetMode(gin.ReleaseMode)
		}

		r := gin.Default()
		group := r.Group("api/v2")
		authGroup := r.Group("api/v2")
		// The middlewares must be registered before controllers register
		middlewares.Register(group, authGroup)

		controllers.Register(group, authGroup)

		address := "0.0.0.0:8080"
		log.Info("Start WebApplication through gin-gonic/gin ", address)

		r.Run(address)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
