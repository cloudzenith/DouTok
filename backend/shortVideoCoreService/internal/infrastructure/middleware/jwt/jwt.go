package jwt

import (
	"context"
	"fmt"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/conf"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	jwt5 "github.com/golang-jwt/jwt/v5"
	"time"
)

func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	whiteList["/shortVideoCoreService.api.v1.UserService/Register"] = struct{}{}
	whiteList["/shortVideoCoreService.api.v1.UserService/Login"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

func Jwt(c *conf.Auth) middleware.Middleware {
	return selector.Server(
		jwt.Server(func(token *jwt5.Token) (interface{}, error) {
			return []byte(c.JWT.AccessSecret), nil
		}, jwt.WithSigningMethod(jwt5.SigningMethodHS256), jwt.WithClaims(func() jwt5.Claims {
			return &jwt5.MapClaims{}
		})),
	).
		Match(NewWhiteListMatcher()).
		Build()
}

func GenerateToken(userId int64, config *conf.Auth) (string, error) {
	hours := time.Duration(config.JWT.AccessExpire)
	token := jwt5.NewWithClaims(jwt5.SigningMethodHS256, jwt5.MapClaims{
		"user_id": userId,
		"exp":     time.Now().UTC().Add(time.Hour * hours).Unix(),
	})
	signedToken, err := token.SignedString([]byte(config.JWT.AccessSecret))
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

	switch claimsAny := claimsAny.(type) {
	case *jwt5.MapClaims:
		claims = claimsAny
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
