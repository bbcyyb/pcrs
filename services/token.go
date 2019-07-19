package services

import (
	"github.com/bbcyyb/pcrs/pkg/jwt"
)

type Token struct {
	jwt *jwt.JWT
}

func NewToken(jwt *jwt.JWT) *Token {
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
