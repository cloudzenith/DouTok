package main

import (
	"context"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/cloudzenith/DouTok/backend/gopkgs/components/mysqlx"
	"github.com/cloudzenith/DouTok/backend/gopkgs/components/redisx"
	"github.com/cloudzenith/DouTok/backend/gopkgs/components/rmqconsumerx"
	"github.com/cloudzenith/DouTok/backend/gopkgs/components/rmqproducerx"
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

func initHttpServer() func(configValue interface{}) *http.Server {
	return func(configValue interface{}) *http.Server {
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

func initGrpcServer() func(configValue interface{}) *grpc.Server {
	return func(configValue interface{}) *grpc.Server {
		srv := grpc.NewServer(
			grpc.Address(":9000"),
		)

		api.RegisterTestServiceServer(srv, application.Application{})
		return srv
	}
}

type TestMessage struct {
	Content string `json:"content"`
	Index   int    `json:"index"`
}

func main() {
	launcher.New(
		launcher.WithConfigValue(&struct{}{}),
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

			consumerx := rmqconsumerx.GetConsumer[*TestMessage](context.Background(), "test")
			err := consumerx.Subscribe(func(ctx context.Context, message *TestMessage, ext *primitive.MessageExt) (consumer.ConsumeResult, error) {
				log.Infow(
					"msg", "received message!",
					"msg_id", ext.MsgId,
					"offset_msg_id", ext.OffsetMsgId,
					"content", message.Content,
					"index", message.Index,
				)
				return consumer.ConsumeSuccess, nil
			})

			if err != nil {
				panic(err)
			}

			producerx := rmqproducerx.GetProducer[*TestMessage](context.Background(), "test")
			for i := 0; i < 10; i++ {
				msg := &TestMessage{
					Content: "hello world - " + string(rune(i)),
					Index:   i,
				}

				_, err := producerx.SendSync(context.Background(), msg)
				if err != nil {
					panic(err)
				}
			}
		}),
	).Run()
}
