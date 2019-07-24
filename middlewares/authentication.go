package middlewares

import (
	"errors"
	"net/http"
	"time"

	"github.com/bbcyyb/pcrs/common"
	pkgJ "github.com/bbcyyb/pcrs/pkg/jwt"
	"github.com/bbcyyb/pcrs/pkg/logger"
	"github.com/gin-gonic/gin"
)

func Authentication(parser pkgJ.IJWTParser) gin.HandlerFunc {
	logger.Log.Debug("Register middleware Authentication")
	return func(c *gin.Context) {
		if err := verify(c, parser); err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
		}

		c.Next()
	}
}

func verify(c *gin.Context, jwt pkgJ.IJWTParser) (err error) {
	logger.Log.Info("Start to JWT Authenticate.")
	var code common.Code
	jwt.SetJwtSecret([]byte("DELLEMC"))

	code = common.SUCCESS
	token := c.GetHeader("X-Authorization")
	if token == "" {
		code = common.ERROR_AUTHT_CHECK_TOKEN_MISS
	} else {
		if claims, err := jwt.Parse(token); err != nil {
			code = common.ERROR_AUTHT_CHECK_TOKEN_FAIL
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = common.ERROR_AUTHT_CHECK_TOKEN_TIMEOUT
		} else {
			c.Set("claims", claims)
		}
	}

	if code != common.SUCCESS {
		msg := common.GetCodeMessage(code)

		err = errors.New(msg)
		logger.Log.Error("JWT verification failed, error message: ", err)
	}

	return
}
