package controllers

import (
	"github.com/bbcyyb/pcrs/infra/log"
	"github.com/gin-gonic/gin"

	"net/http"
)

type Miscellaneous struct {
}

func NewMiscllaneous() *Miscellaneous {
	return &Miscellaneous{}
}

func (misc *Miscellaneous) Test(c *gin.Context) {
	log.Debug("Start to run test action.")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}
