package services

import (
	"github.com/bbcyyb/pcrs/pkg/jwt"
)

type Token struct {
	jwt *jwt.Jwt
}

func NewToken(jwt *jwt.Jwt) *Token {
	return &Token{
		jwt: jwt,
	}
}

type TokenHandler interface {
	GenerateToken(*jwt.Claims) (string, error)
}

func (t *Token) GenerateToken(claims *jwt.Claims) (string, error) {
	return t.jwt.Generate(claims)
}
