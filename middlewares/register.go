package middlewares

import (
	"github.com/bbcyyb/pcrs/pkg/log"
	"github.com/gin-gonic/gin"
)

func Register(group *gin.RouterGroup, authGroup *gin.RouterGroup) {
	log.Info("Register middlewares")

	authGroup.Use(ErrorHandler())
	group.Use(ErrorHandler())

	if C.AuthTEnable {
		authGroup.Use(Authentication())
	}

	if C.AuthREnable {
		authGroup.Use(Authorization())
	}

	if C.LogEnable {
		authGroup.Use(Log())
		group.Use(Log())
	}
}
