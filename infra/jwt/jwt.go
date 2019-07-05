package jwt

import (
	"errors"
	"github.com/bbcyyb/pcrs/common"
	jg "github.com/dgrijalva/jwt-go"
	"time"
)

const (
	timeOffset      time.Duration = 7 * 24 * time.Hour
	debugTimeOffset time.Duration = 10 * 365 * 24 * time.Hour
)

var (
	jwtSecret []byte = []byte("DELLEMC")
)

type Claims struct {
	Id       int             `json:"id"`
	RsaId    string          `json:"rid"`
	UserName string          `json:"un"`
	Email    string          `json:"em"`
	Role     common.RoleType `json:"ur"`
	IsDebug  int             `json:"de"`
	jg.StandardClaims
}

func GenerateToken(claims *Claims) (token string, err error) {
	/*
		nowTime := time.Now()
		var expireTime time.Time
		if claims.IsDebug == 0 {
			expireTime = nowTime.Add(timeOffset)
		} else {
			expireTime = nowTime.Add(debugTimeOffset)
		}

		claims.ExpiresAt = expireTime.Unix()
		claims.Issuer = "powercalculator"
	*/

	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("Unknow panic")
			}
		}
	}()

	tokenClaims := jg.NewWithClaims(jg.SigningMethodHS256, *claims)
	token, err = tokenClaims.SignedString(jwtSecret)

	return
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jg.ParseWithClaims(token, &Claims{}, func(token *jg.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
