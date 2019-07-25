package jwt

import (
	"errors"
	"github.com/bbcyyb/pcrs/common"
	jg "github.com/dgrijalva/jwt-go"
)

var (
	defaultJwtSecret []byte = []byte("DELLEMC")
)

type Jwt struct {
	jwtSecret []byte
}

type IJwtParser interface {
	Parse(string) (*Claims, error)
	SetJwtSecret([]byte)
}

type IJwtGenerator interface {
	Generate(*Claims) (string, error)
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

func NewJwt() *Jwt {
	return &Jwt{}
}

func (j *Jwt) Generate(claims *Claims) (token string, err error) {

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

func (j *Jwt) Parse(token string) (*Claims, error) {
	tokenClaims, err := jg.ParseWithClaims(token, &Claims{}, func(token *jg.Token) (interface{}, error) {
		return j.jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			if claims.ExpiresAt > 1000000000000 {
				// if expire time length is 13, the first 10 bits need to be intercepted.
				claims.ExpiresAt /= 1000
			}

			return claims, nil
		}
	}

	return nil, err
}

func (j *Jwt) SetJwtSecret(jwtSecret []byte) {
	j.jwtSecret = jwtSecret
}
