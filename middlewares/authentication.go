package middlewares

import (
	"errors"
	"net/http"
	"time"

	"github.com/bbcyyb/pcrs/common"
	pkgJ "github.com/bbcyyb/pcrs/pkg/jwt"
	"github.com/bbcyyb/pcrs/pkg/log"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	log.Debug("Register middleware JWTAuth")
	return func(c *gin.Context) {
		if err := verify(c); err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
		}

		c.Next()
	}
}

func verify(c *gin.Context) (err error) {
	log.Info("Start to JWT Authenticate.")
	var code common.Code
	jwt := pkgJ.NewJWT()
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
		log.Error("JWT verification failed, error message: ", err)
	}

	return
}
