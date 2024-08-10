package middleware

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/conf"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	jwt5 "github.com/golang-jwt/jwt/v5"
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

func Jwt(c *conf.Common) middleware.Middleware {
	return selector.Server(
		jwt.Server(func(token *jwt5.Token) (interface{}, error) {
			return []byte(c.Auth.Secret), nil
		}, jwt.WithSigningMethod(jwt5.SigningMethodHS256), jwt.WithClaims(func() jwt5.Claims {
			return &jwt5.MapClaims{}
		})),
	).
		Match(NewWhiteListMatcher()).
		Build()
}
