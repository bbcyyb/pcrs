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
	defaultJwtSecret []byte = []byte("DEFAULT")
)

type JWT struct {
	jwtSecret []byte
}

type Claims struct {
	Id       int             `json:"id"`
	RsaId    string          `json:"rid"`
	UserName string          `json:"un"`
	Email    string          `json:"em"`
	Role     common.RoleEnum `json:"ur"`
	IsDebug  int             `json:"de"`
	jg.StandardClaims
}

type JWTHandler interface {
	Generate(*Claims) (string, error)
	Parse(string) (*Claims, error)
}

func NewJWT() *JWT {
	return &JWT{}
}

func (j *JWT) Generate(claims *Claims) (token string, err error) {
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
	token, err = tokenClaims.SignedString(j.jwtSecret)

	return
}

func (j *JWT) Parse(token string) (*Claims, error) {
	tokenClaims, err := jg.ParseWithClaims(token, &Claims{}, func(token *jg.Token) (interface{}, error) {
		return j.jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

func (j *JWT) SetJwtSecret(jwtSecret []byte) {
	j.jwtSecret = jwtSecret
}
