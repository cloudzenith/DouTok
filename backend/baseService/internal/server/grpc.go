package server

import (
	"github.com/cloudzenith/DouTok/backend/baseService/api"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/middlewares"
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

	warmUp(params)

	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			metadata.Server(),
			tracing.Server(),
			middlewares.TraceIdInjector(),
			middlewares.SpanIdInjector(),
			middlewares.RequestMonitor(),
			middlewares.ProtobufValidator(),
		),
	}

	if params.addr != "" {
		opts = append(opts, grpc.Address(params.addr))
	}

	srv := grpc.NewServer(opts...)

	api.RegisterAccountServiceServer(srv, initAccountApplication())
	api.RegisterAuthServiceServer(srv, initAuthApplication())
	api.RegisterPostServiceServer(srv, initPostApplication())
	api.RegisterFileServiceServer(srv, initFileApplication(params.fileTableShardingConfig))
	return srv
}
