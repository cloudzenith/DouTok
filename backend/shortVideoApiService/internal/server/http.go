package server

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/api/svapi"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/middlewares"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/utils/claims"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwt5 "github.com/golang-jwt/jwt/v5"
)

func TokenParseWhiteList() selector.MatchFunc {
	whileList := make(map[string]struct{})
	whileList["/user/code"] = struct{}{}
	whileList["/user/login"] = struct{}{}
	whileList["/user/register"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whileList[operation]; ok {
			return true
		}

		return false
	}
}

func NewHttpServer() *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			middlewares.ResponseWrapper(),
			selector.Server(
				jwt.Server(
					func(token *jwt5.Token) (interface{}, error) {
						return "token", nil
					},
					jwt.WithClaims(func() jwt5.Claims {
						return &claims.Claims{}
					}),
				),
			).Match(TokenParseWhiteList()).Build(),
		),
	}

	srv := http.NewServer(opts...)

	svapi.RegisterUserServiceHTTPServer(srv, initUserApp())
	return srv
}
