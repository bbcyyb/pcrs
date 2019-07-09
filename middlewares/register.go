package middlewares

import (
	"github.com/bbcyyb/pcrs/infra/log"
	"github.com/gin-gonic/gin"
)

func Register(api *gin.Engine) {
	//func Register(api *gin.RouterGroup) {
	log.Info("Register middlewares")

	api.Use(func(c *gin.Context) {
		log.Info("111111111")
		c.Next()
		log.Info("222222222")
	})

	if C.JWTEnable {
		api.Use(JWTAuth())
	}

	if C.AuthEnable {
		api.Use(Authorization())
	}

	if C.LogEnable {
		api.Use(Log())
	}
}
