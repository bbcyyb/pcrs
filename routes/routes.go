package routes

import (
	"github.com/bbcyyb/pcrs/controllers"
	"github.com/bbcyyb/pcrs/infra/log"
	"github.com/gin-gonic/gin"
)

func Register(api *gin.RouterGroup) {
	log.Info("Register restful service route handler")
	registerMiscellaneous(api)
}

func registerMiscellaneous(api *gin.RouterGroup) {
	log.Debug("Register Miscellaneous route handler")
	g := api.Group("/tests")
	g.GET("/test", controllers.Test)

}
