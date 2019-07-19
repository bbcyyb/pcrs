package middlewares

import (
	"net/http"

	. "github.com/bbcyyb/pcrs/common"
	"github.com/bbcyyb/pcrs/pkg/log"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	log.Debug("Register middleware ResponseHandler")
	return func(c *gin.Context) {
		c.Next()

		status := c.Writer.Status()
		if status > http.StatusBadRequest {
			Respond(c, status, nil, c.Errors.Last())
		}
	}
}
