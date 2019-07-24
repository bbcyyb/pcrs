package middlewares

import (
	"net/http"

	"github.com/bbcyyb/pcrs/common"
	"github.com/bbcyyb/pcrs/pkg/logger"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	logger.Log.Debug("Register middleware ResponseHandler")
	return func(c *gin.Context) {
		c.Next()

		status := c.Writer.Status()
		if status > http.StatusBadRequest {
			common.Respond(c, status, nil, c.Errors.Last())
		}
	}
}
