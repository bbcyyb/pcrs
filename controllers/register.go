package controllers

import (
	"github.com/bbcyyb/pcrs/infra/jwt"
	"github.com/bbcyyb/pcrs/infra/log"
	"github.com/bbcyyb/pcrs/services"
	"github.com/gin-gonic/gin"
)

func Register(group *gin.RouterGroup, authGroup *gin.RouterGroup) {
	log.Info("Register restful service route handler")

	registerMiscellaneous(group)
}

func registerMiscellaneous(g *gin.RouterGroup) {
	log.Debug("Register Miscellaneous route handler")
	j := jwt.NewJWT()
	j.SetJwtSecret([]byte("DELLEMC"))
	token := services.NewToken(j)
	misc := NewMiscllaneous(token)

	miscG := g.Group("/miscs")
	miscG.GET("/test", misc.Test)
	miscG.POST("/token", misc.GenerateTestToken)
}
