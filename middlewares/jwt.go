package middlewares

import (
	"github.com/bbcyyb/pcrs/infra/log"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

var jwtSecret = []byte(viper.GetString("jwtSecret"))
var timeOffset = viper.GetInt("timeOffset")

type Claims struct {
	User
	jwt.StandardClaims
}

type User struct {
	Id       string `json:"id"`
	RsaId    string `json:"rsaid"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Role     int    `json:"role"`
}

func GenerateToken(user User) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(timeOffset) * time.Hour)

	claims := Claims{
		user,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "DELLEMC",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Info("======> start to jwt auth.")
		var code int
		var data interface{}

		code = 200
		token := c.Query("token")
		if token == "" {
			code = 404
		} else {
			claims, err := ParseToken(token)
			if err != nil {
				code = 401
			} else if claims == nil {
				code = 403
			}
		}

		if code != 200 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  "unauthorized",
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
