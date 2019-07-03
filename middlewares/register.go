package middlewares

import (
	"github.com/bbcyyb/pcrs/infra/log"
	"github.com/bbcyyb/pcrs/middlewares/jwt"
	"github.com/gin-gonic/gin"
)

func Register(api *gin.RouterGroup) {
	log.Info("Register middlewares")

	if C.JWT {
		api.Use(jwt.JWT())
	}
}
