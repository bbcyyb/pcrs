package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func OK(c *gin.Context, data interface{}) {
	Respond(c, http.StatusOK, data, nil)
}

func Created(c *gin.Context, data interface{}) {
	Respond(c, http.StatusCreated, data, nil)
}

func Accepted(c *gin.Context) {
	Respond(c, http.StatusAccepted, nil, nil)
}

func NoContent(c *gin.Context) {
	Respond(c, http.StatusNoContent, nil, nil)
}

func BadRequest(c *gin.Context, err error) {
	c.AbortWithError(http.StatusBadRequest, err)
	Respond(c, http.StatusBadRequest, nil, nil)
}

func Unauthorized(c *gin.Context, err error) {
	c.AbortWithError(http.StatusUnauthorized, err)
	Respond(c, http.StatusUnauthorized, nil, nil)
}

func Forbidden(c *gin.Context, err error) {
	c.AbortWithError(http.StatusForbidden, err)
	Respond(c, http.StatusForbidden, nil, nil)
}

func NotFound(c *gin.Context, err error) {
	c.AbortWithError(http.StatusNotFound, err)
	Respond(c, http.StatusNotFound, nil, nil)
}

func InternalServerError(c *gin.Context, err error) {
	c.AbortWithError(http.StatusInternalServerError, err)
	Respond(c, http.StatusInternalServerError, nil, nil)
}

func Respond(c *gin.Context, status int, data interface{}, err error) {

	resp := Response{
		Code:    status,
		Message: err,
		Data:    data,
	}

	if gin.Mode() == gin.ReleaseMode {
		c.JSON(status, resp)
	} else {
		c.IndentedJSON(status, resp)
	}
}
