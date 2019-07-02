package cmd

import (
	"github.com/bbcyyb/pcrs/infra/log"
	"github.com/bbcyyb/pcrs/routes"
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

		log.Debug("Declare Group api/v2")
		api := r.Group("api/v2")
		routes.Register(api)

		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		address := "0.0.0.0:" + strconv.Itoa(C.Port)
		log.Info("Start WebApplication through gin-gonic/gin ", address)

		r.Run(address)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
