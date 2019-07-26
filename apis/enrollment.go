package apis

import (
	"github.com/bbcyyb/pcrs/pkg/logger"
	"github.com/gin-gonic/gin"
)

func Enroll(group *gin.RouterGroup, authGroup *gin.RouterGroup) {
	logger.Log.Info("Enroll restful service route handler")

	enrollMisc(group, authGroup)
}

func enrollMisc(group *gin.RouterGroup, authGroup *gin.RouterGroup) {
	controller, _ := initializeMisc()

	miscG := group.Group("/miscs")
	miscG.POST("tokens", controller.GenerateTestToken)

	miscAg := authGroup.Group("/miscs")
	miscAg.GET("/test", controller.Test)
}
