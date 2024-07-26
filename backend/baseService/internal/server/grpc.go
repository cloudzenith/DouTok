package server

import (
	"github.com/cloudzenith/DouTok/backend/baseService/api"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/middlewares"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/server/authappproviders"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

func NewGRPCServer(options ...Option) *grpc.Server {
	params := &Params{}
	for _, option := range options {
		option(params)
	}

	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			metadata.Server(),
			tracing.Server(),
			middlewares.TraceIdInjector(),
			middlewares.SpanIdInjector(),
			middlewares.RequestMonitor(),
		),
	}
	opts = append(opts, grpc.Address(params.addr))
	srv := grpc.NewServer(opts...)

	api.RegisterAccountServiceServer(srv, initAccountApplication())
	api.RegisterAuthServiceServer(srv, initAuthApplication(authappproviders.RedisDsn(params.redisDsn), authappproviders.RedisPassword(params.redisPassword)))
	return srv
}
