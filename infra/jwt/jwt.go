package jwt

import (
	"github.com/bbcyyb/pcrs/common"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret []byte

type Claims struct {
	Id          int             `json:"id"`
	RsaId       string          `json:"rid"`
	UserName    string          `json:"un"`
	Email       string          `json:"em"`
	Role        common.RuleType `json:"ur"`
	ExpiredTime float32         `json:"exp"`
	IsDebug     int             `json:"de"`
	jwt.StandardClaims
}

func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{}
}
