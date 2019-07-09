package cmd

import (
	"github.com/bbcyyb/pcrs/controllers"
	"github.com/bbcyyb/pcrs/infra/log"
	"github.com/bbcyyb/pcrs/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"strconv"
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
		api := r.Group("api/v2")
		controllers.Register(api)
		middlewares.Register(r)

		address := "0.0.0.0:" + strconv.Itoa(C.Port)
		log.Info("Start WebApplication through gin-gonic/gin ", address)

		r.Run(address)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
