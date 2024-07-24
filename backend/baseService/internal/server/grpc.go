package server

import (
	"github.com/cloudzenith/DouTok/backend/baseService/api"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/middlewares"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

func NewGRPCServer(addr string) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			metadata.Server(),
			logging.Server(log.GetLogger()),
			tracing.Server(),
			middlewares.TraceIdInjector(),
			middlewares.SpanIdInjector(),
		),
	}
	opts = append(opts, grpc.Address(addr))
	srv := grpc.NewServer(opts...)

	api.RegisterAccountServiceServer(srv, initAccountApplication())
	return srv
}
