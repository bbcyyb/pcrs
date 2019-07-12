package controllers

import (
	. "github.com/bbcyyb/pcrs/common"
	"github.com/bbcyyb/pcrs/infra/jwt"
	"github.com/bbcyyb/pcrs/services"
	"github.com/gin-gonic/gin"

	"time"
)

const (
	timeOffset      time.Duration = 7 * 24 * time.Hour
	debugTimeOffset time.Duration = 10 * 365 * 24 * time.Hour
)

type Miscellaneous struct {
	tokenHandler services.TokenHandler
}

func NewMiscllaneous(tokenHandler services.TokenHandler) *Miscellaneous {
	return &Miscellaneous{
		tokenHandler: tokenHandler,
	}
}

func (misc *Miscellaneous) Test(c *gin.Context) {
	OK(c, gin.H{
		"message": "Hello World!",
	})
}

func (misc *Miscellaneous) GenerateTestToken(c *gin.Context) {
	var claims jwt.Claims
	if err := c.BindJSON(&claims); err != nil {
		BadRequest(c, err)
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
		InternalServerError(c, err)
	} else {
		OK(c, gin.H{
			"token": token,
		})
	}
}
