package middlewares

import (
	pkgA "github.com/bbcyyb/pcrs/pkg/authorizer"
	pkgJ "github.com/bbcyyb/pcrs/pkg/jwt"
	"github.com/bbcyyb/pcrs/pkg/log"
	"github.com/gin-gonic/gin"
)

func Register(group *gin.RouterGroup, authGroup *gin.RouterGroup) {
	log.Info("Register middlewares")

	authGroup.Use(ErrorHandler())
	group.Use(ErrorHandler())

	if C.AuthTEnable {
		jwt := pkgJ.NewJWT()
		authGroup.Use(Authentication(jwt))
	}

	if C.AuthREnable {
		authorizer := pkgA.NewBasicAuthorizer()
		authGroup.Use(Authorization(authorizer))
	}

	if C.LogEnable {
		authGroup.Use(Log())
		group.Use(Log())
	}
}
