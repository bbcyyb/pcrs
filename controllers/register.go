package controllers

import (
	"github.com/bbcyyb/pcrs/infra/log"
	"github.com/gin-gonic/gin"
)

func Register(api *gin.RouterGroup) {
	log.Info("Register restful service route handler")
	registerMiscellaneous(api)
}

func registerMiscellaneous(api *gin.RouterGroup) {
	log.Debug("Register Miscellaneous route handler")
	misc := NewMiscllaneous()
	g := api.Group("/miscs")
	g.GET("/test", misc.Test)

}
