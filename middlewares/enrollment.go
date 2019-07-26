package middlewares

import (
	"github.com/bbcyyb/pcrs/conf"
	pkgA "github.com/bbcyyb/pcrs/pkg/authorizer"
	pkgJ "github.com/bbcyyb/pcrs/pkg/jwt"
	"github.com/bbcyyb/pcrs/pkg/logger"
	"github.com/gin-gonic/gin"
)

func Enroll(group *gin.RouterGroup, authGroup *gin.RouterGroup) {
	logger.Log.Info("Register middlewares")

	authGroup.Use(ErrorHandler())
	group.Use(ErrorHandler())

	if conf.C.Middleware.Authentication.Enable {
		jwt := pkgJ.NewJwtParser()
		authGroup.Use(Authentication(jwt))
	}

	if conf.C.Middleware.Authorization.Enable {
		authorizer := pkgA.NewBasicAuthorizer()
		authGroup.Use(Authorization(authorizer))
	}

	if conf.C.Middleware.ErrorHandler.Enable {
		authGroup.Use(ErrorHandler())
		group.Use(ErrorHandler())
	}
}
