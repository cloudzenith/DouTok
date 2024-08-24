package claims

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwt5 "github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt5.RegisteredClaims
	UserId int64 `json:"user_id"`
}

func New(userId int64) *Claims {
	return &Claims{
		UserId: userId,
	}
}

func GetUserId(ctx context.Context) (int64, error) {
	anyClaims, ok := jwt.FromContext(ctx)
	if !ok {
		return 0, errors.New("no claims in context")
	}

	claims, ok := anyClaims.(*Claims)
	if !ok {
		return 0, errors.New("claims type error")
	}

	return claims.UserId, nil
}
