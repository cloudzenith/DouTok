package server

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/api/svapi"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/utils/claims"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwt5 "github.com/golang-jwt/jwt/v5"
)

func TokenParseWhiteList() selector.MatchFunc {
	whileList := make(map[string]struct{})
	whileList["/svapi.UserService/GetVerificationCode"] = struct{}{}
	whileList["/svapi.UserService/Register"] = struct{}{}
	whileList["/svapi.UserService/Login"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whileList[operation]; ok {
			return false
		}

		return true
	}
}

func NewHttpServer() *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			selector.Server(
				jwt.Server(
					func(token *jwt5.Token) (interface{}, error) {
						return []byte("token"), nil
					},
					jwt.WithClaims(func() jwt5.Claims {
						return &claims.Claims{}
					}),
				),
			).Match(TokenParseWhiteList()).Build(),
			// 这个中间件包装返回值之后，会导致返回值的类型不匹配，所以暂时注释掉
			// 详见断言：backend/shortVideoApiService/api/svapi/user_http.pb.go:67
			//middlewares.ResponseWrapper(),
		),
		http.Address("0.0.0.0:22000"),
	}

	srv := http.NewServer(opts...)

	svapi.RegisterUserServiceHTTPServer(srv, initUserApp())
	return srv
}
