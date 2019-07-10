package controllers

import (
	"github.com/bbcyyb/pcrs/infra/jwt"
	"github.com/bbcyyb/pcrs/services"
	"github.com/gin-gonic/gin"

	"net/http"
	"time"
)

const (
	timeOffset      time.Duration = 7 * 24 * time.Hour
	debugTimeOffset time.Duration = 10 * 365 * 24 * time.Hour
)

type Miscellaneous struct {
	Base
	tokenHandler services.TokenHandler
}

func NewMiscllaneous(tokenHandler services.TokenHandler) *Miscellaneous {
	return &Miscellaneous{
		tokenHandler: tokenHandler,
	}
}

func (misc *Miscellaneous) Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}

func (misc *Miscellaneous) GenerateTestToken(c *gin.Context) {
	var claims jwt.Claims
	if err := c.BindJSON(&claims); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	nowTime := time.Now()
	var expireTime time.Time
	if claims.IsDebug == 0 {
		expireTime = nowTime.Add(timeOffset)
	} else {
		expireTime = nowTime.Add(debugTimeOffset)
	}

	claims.ExpiresAt = expireTime.Unix()
	claims.Issuer = "powercalculator"

	token, err := misc.tokenHandler.GenerateToken(&claims)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": token,
		})
	}
}
