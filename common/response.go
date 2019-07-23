package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func OK(c *gin.Context, data interface{}) {
	Respond(c, http.StatusOK, data, nil)
}

func Created(c *gin.Context, data interface{}) {
	Respond(c, http.StatusCreated, data, nil)
}

func Accepted(c *gin.Context, data interface{}) {
	Respond(c, http.StatusAccepted, data, nil)
}

func NoContent(c *gin.Context) {
	Respond(c, http.StatusNoContent, nil, nil)
}

func BadRequest(c *gin.Context, err error) {
	c.AbortWithError(http.StatusBadRequest, err)
	Respond(c, http.StatusBadRequest, nil, err)
}

func Unauthorized(c *gin.Context, err error) {
	c.AbortWithError(http.StatusUnauthorized, err)
	Respond(c, http.StatusUnauthorized, nil, err)
}

func Forbidden(c *gin.Context, err error) {
	c.AbortWithError(http.StatusForbidden, err)
	Respond(c, http.StatusForbidden, nil, err)
}

func NotFound(c *gin.Context, err error) {
	c.AbortWithError(http.StatusNotFound, err)
	Respond(c, http.StatusNotFound, nil, err)
}

func InternalServerError(c *gin.Context, err error) {
	c.AbortWithError(http.StatusInternalServerError, err)
	Respond(c, http.StatusInternalServerError, nil, err)
}

func Respond(c *gin.Context, status int, data interface{}, err error) {
	resp := Response{
		Code:    status,
		Message: "",
		Data:    data,
	}

	if err != nil {
		resp.Message = err.Error()
	}

	if gin.Mode() == gin.ReleaseMode {
		c.JSON(status, resp)
	} else {
		c.IndentedJSON(status, resp)
	}
}
