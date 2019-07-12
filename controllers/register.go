package controllers

import (
	"github.com/bbcyyb/pcrs/infra/jwt"
	"github.com/bbcyyb/pcrs/infra/log"
	"github.com/bbcyyb/pcrs/services"
	"github.com/gin-gonic/gin"
)

func Register(group *gin.RouterGroup, authGroup *gin.RouterGroup) {
	log.Info("Register restful service route handler")

	registerMiscellaneous(group, authGroup)
}

func registerMiscellaneous(group *gin.RouterGroup, authGroup *gin.RouterGroup) {
	log.Debug("Register Miscellaneous route handler")
	j := jwt.NewJWT()
	j.SetJwtSecret([]byte("DELLEMC"))
	tokenSvc := services.NewToken(j)
	miscController := NewMiscllaneous(tokenSvc)

	miscG := group.Group("/miscs")
	miscG.POST("/tokens", miscController.GenerateTestToken)

	miscAg := authGroup.Group("/miscs")
	miscAg.GET("/test", miscController.Test)
}
