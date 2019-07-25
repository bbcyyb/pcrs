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
	initializeMisc()
}
