package auth

import (
	"context"
	"fmt"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/conf"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwt5 "github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateToken(userId int64, config *conf.Common) (string, error) {
	hours := time.Duration(config.Auth.Expire)
	token := jwt5.NewWithClaims(jwt5.SigningMethodHS256, jwt5.MapClaims{
		"user_id": userId,
		"exp":     time.Now().UTC().Add(time.Hour * hours).Unix(),
	})
	signedToken, err := token.SignedString([]byte(config.Auth.Secret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func GetLoginUser(ctx context.Context) (int64, error) {
	// 从 context 中获取验证后的 Token
	claimsAny, ok := jwt.FromContext(ctx)
	if !ok {
		return 0, fmt.Errorf("token not found in context")
	}

	// 断言 claims 是 jwt.MapClaims 类型
	var claims *jwt5.MapClaims

	switch claimsAny.(type) {
	case *jwt5.MapClaims:
		claims = claimsAny.(*jwt5.MapClaims)
	default:
		fmt.Println("claims is not of type jwt.MapClaims")
		return 0, fmt.Errorf("claims is not of type jwt.MapClaims")
	}

	// 从 claims 中提取 user_id
	userId, exists := (*claims)["user_id"]
	if !exists {
		return 0, fmt.Errorf("user_id not found in claims")
	}
	return int64(userId.(float64)), nil
}
