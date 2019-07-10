package middlewares

import (
	"github.com/bbcyyb/pcrs/infra/log"
	"github.com/gin-gonic/gin"
)

func Register(group *gin.RouterGroup, authGroup *gin.RouterGroup) {
	log.Info("Register middlewares")

	if C.JWTEnable {
		authGroup.Use(JWTAuth())
	}

	if C.AuthEnable {
		authGroup.Use(Authorization())
	}

	if C.LogEnable {
		authGroup.Use(Log())
		group.Use(Log())
	}
}
