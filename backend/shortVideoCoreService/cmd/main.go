package main

import (
	"flag"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/conf"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/data"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/userdomain"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/pkg/db"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/pkg/utils"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/server"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/service/thirdparty"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/service/userservice"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"os"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// confFlag is the config flag.
	confFlag string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&confFlag, "conf", "./configs/config.yaml", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
	)
}

func InitApp(config *conf.Config, logger log.Logger) (*kratos.App, error) {
	dbClient, err := db.NewDBClient(config.Database)
	if err != nil {
		return nil, err
	}
	snowflake, err := utils.NewSnowflakeNode(config.Common)
	if err != nil {
		return nil, err
	}
	userRepo := data.NewUserRepo(dbClient, logger)
	baseService, err := thirdparty.NewBaseService(config.ThirdParty, logger)
	if err != nil {
		return nil, err
	}
	userUsecase := userdomain.NewUserUsecase(
		config, snowflake, baseService.AccountServiceClient, userRepo, dbClient, logger,
	)
	userService := userservice.NewUserService(config, logger, userUsecase)
	grpcServer := server.NewGRPCServer(config, userService, logger)
	httpServer := server.NewHTTPServer(config, userService, logger)
	return newApp(logger, grpcServer, httpServer), err
}

func main() {
	flag.Parse()
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	c := config.New(
		config.WithSource(
			file.NewSource(confFlag),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Config

	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	app, err := InitApp(&bc, logger)
	if err != nil {
		panic(err)
	}

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
