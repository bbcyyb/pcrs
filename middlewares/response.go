package middlewares

import (
	"net/http"

	"github.com/bbcyyb/pcrs/infra/log"
	"github.com/gin-gonic/gin"
)

func ResponseHandler() gin.HandlerFunc {
	log.Debug("Register middleware ResponseHandler")
	return func(c *gin.Context) {
		c.Next()

		status := c.Writer.Status()
		switch status {
		case http.StatusOK:
			log.Info("StatusOK")
		}
	}
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func json200(c *gin.Context, data interface{}) {
	respond(c, http.StatusOK, "", data)
}

func json201(c *gin.Context, data interface{}) {
	respond(c, http.StatusCreated, "", data)
}

func json202(c *gin.Context) {
	respond(c, http.StatusAccepted, "", nil)
}

func json204(c *gin.Context) {
	respond(c, http.StatusNoContent, "", nil)
}

func json400(c *gin.Context, message string) {
	respond(c, http.StatusBadRequest, message, nil)
}

func json401(c *gin.Context, message string) {
	respond(c, http.StatusUnauthorized, message, nil)
}

func json403(c *gin.Context, message string) {
	respond(c, http.StatusForbidden, message, nil)
}

func json404(c *gin.Context, message string) {
	respond(c, http.StatusNotFound, message, nil)
}

func json500(c *gin.Context, message string) {
	respond(c, http.StatusInternalServerError, message, nil)
}

func respond(c *gin.Context, status int, message string, data interface{}) {
	resp := Response{
		Code:    status,
		Message: message,
		Data:    data,
	}

	if gin.Mode() == gin.ReleaseMode {
		c.JSON(status, resp)
	} else {
		c.IndentedJSON(status, resp)
	}
}
