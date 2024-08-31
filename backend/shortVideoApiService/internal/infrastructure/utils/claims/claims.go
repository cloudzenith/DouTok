package claims

import (
	"context"
	"errors"
	"fmt"
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

func GenerateToken(claim *Claims) (string, error) {
	token := jwt5.NewWithClaims(jwt5.SigningMethodHS256, claim)
	tokenString, err := token.SignedString([]byte("token"))
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}
	return tokenString, nil
}
