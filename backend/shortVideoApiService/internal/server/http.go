package server

import (
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/api"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/applications/userapp"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/middlewares"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func NewGinServer() *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			middlewares.ResponseWrapper(),
		),
	}

	srv := http.NewServer(opts...)

	api.RegisterUserServiceHTTPServer(srv, userapp.New())
	return srv
}
