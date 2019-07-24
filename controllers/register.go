package controllers

import (
	"github.com/bbcyyb/pcrs/pkg/jwt"
	"github.com/bbcyyb/pcrs/pkg/logger"
	"github.com/bbcyyb/pcrs/services"
	"github.com/gin-gonic/gin"
)

func Register(group *gin.RouterGroup, authGroup *gin.RouterGroup) {
	logger.Log.Info("Register restful service route handler")

	registerMiscellaneous(group, authGroup)
}

func registerMiscellaneous(group *gin.RouterGroup, authGroup *gin.RouterGroup) {
	logger.Log.Debug("Register Miscellaneous route handler")
	j := jwt.NewJWT()
	j.SetJwtSecret([]byte("DELLEMC"))
	tokenSvc := services.NewToken(j)
	miscController := NewMiscllaneous(tokenSvc)

	miscG := group.Group("/miscs")
	miscG.POST("/tokens", miscController.GenerateTestToken)

	miscAg := authGroup.Group("/miscs")
	miscAg.GET("/test", miscController.Test)
}
