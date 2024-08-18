package main

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/gopkgs/components/mysqlx"
	"github.com/cloudzenith/DouTok/backend/gopkgs/components/redisx"
	"github.com/cloudzenith/DouTok/backend/gopkgs/launcher"
	"github.com/cloudzenith/DouTok/backend/gopkgs/launcher/example/api"
	"github.com/cloudzenith/DouTok/backend/gopkgs/launcher/example/application"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"gopkg.in/yaml.v2"
)

func initHttpServer() func() *http.Server {
	return func() *http.Server {
		srv := http.NewServer(
			http.Address(":8000"),
		)

		redisClient := redisx.GetClient(context.Background())
		if err := redisClient.Ping(context.Background()).Err(); err != nil {
			panic(err)
		} else {
			log.Info("redis could be loaded before register http server")
		}

		mysqlClient := mysqlx.GetDBClient(context.Background())
		db, _ := mysqlClient.DB()
		if err := db.Ping(); err != nil {
			panic(err)
		} else {
			log.Info("mysql could be loaded before register http server")
		}

		api.RegisterTestServiceHTTPServer(srv, application.Application{})
		return srv
	}
}

func initGrpcServer() func() *grpc.Server {
	return func() *grpc.Server {
		srv := grpc.NewServer(
			grpc.Address(":9000"),
		)

		api.RegisterTestServiceServer(srv, application.Application{})
		return srv
	}
}

func main() {
	launcher.New(
		launcher.WithConfigOptions(
			config.WithSource(
				file.NewSource("configs/"),
			),
			config.WithDecoder(func(keyValue *config.KeyValue, m map[string]interface{}) error {
				return yaml.Unmarshal(keyValue.Value, m)
			}),
		),
		launcher.WithHttpServer(initHttpServer()),
		launcher.WithGrpcServer(initGrpcServer()),
		launcher.WithAfterServerStartHandler(func() {
			redisClient := redisx.GetClient(context.Background())
			if err := redisClient.Ping(context.Background()).Err(); err != nil {
				panic(err)
			} else {
				log.Info("redis connected")
			}
		}),
	).Run()
}
