package jwt

import (
	"errors"
	"net/http"
	"time"

	"github.com/bbcyyb/pcrs/common"
	j "github.com/bbcyyb/pcrs/infra/jwt"
	"github.com/bbcyyb/pcrs/infra/log"
	"github.com/gin-gonic/gin"
	//"gin-blog/pkg/util"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := verify(c); err != nil {
			return
		}

		c.Next()
	}
}

func verify(c *gin.Context) (err error) {
	log.Info("Start to JWT Authenticate.")
	var code common.Code
	var data interface{}

	code = common.SUCCESS
	token := c.Query("token")
	if token == "" {
		code = common.INVALID_PARAMS
	} else {
		claims, err := j.ParseToken(token)
		if err != nil {
			code = common.ERROR_AUTH_CHECK_TOKEN_FAIL
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = common.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
		}
	}

	if code != common.SUCCESS {
		msg := common.GetMsg(code)
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": code,
			"msg":  msg,
			"data": data,
		})

		c.Abort()
		err = errors.New(msg)
		log.Debug("JWT verification failed, error message: ", msg)
	}

	return
}
