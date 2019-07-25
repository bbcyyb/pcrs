package misc

import (
	"context"
	"github.com/bbcyyb/pcrs/pkg/jwt"
)

const (
	timeOffset      time.Duration = 7 * 24 * time.Hour
	debugTimeOffset time.Duration = 10 * 365 * 24 * time.Hour
)

type IService interface {
	GenerateToken(context.Context, *jwt.Claims) (string, error)
}

type MiscService struct {
	jwtGenerator jwt.IJwtGenerator
}

func NewMiscService(generator jwt.IJwtGenerator) *MiscService {
	return &MiscService{
		jwtGenerator: generator,
	}
}

func (misc *MiscService) GenerateToken(ctx context.Context, claims *jwt.Claims) (string, error) {
	nowTime := time.Now()
	var expireTime time.Time
	if claims.IsDebug == 0 {
		expireTime = nowTime.Add(timeOffset)
	} else {
		expireTime = nowTime.Add(debugTimeOffset)
	}

	claims.ExpiresAt = expireTime.Unix()
	claims.Issuer = "powercalculator"

	return misc.jwtGenerator.Generate(claims)
}
