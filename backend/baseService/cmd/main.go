package main

import (
	"github.com/cloudzenith/DouTok/backend/baseService/internal/conf"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/dal/query"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/utils"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/server"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var logger log.Logger

func newBaseService(logger log.Logger, gs *grpc.Server) *kratos.App {
	return kratos.New(
		kratos.Logger(logger),
		kratos.Server(gs),
	)
}

func loadConfig() (config.Config, func()) {
	c := config.New(
		config.WithSource(
			file.NewSource("./configs"),
		),
	)

	if err := c.Load(); err != nil {
		panic(err)
	}

	return c, func() {
		c.Close()
	}
}

func newDb(cfg conf.Config) {
	db, err := gorm.Open(
		mysql.Open(cfg.Data.Database.Source),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	query.SetDefault(db)
}

func init() {
	logger = utils.SetJsonLogger()
}

func main() {
	cfg, closeCfg := loadConfig()
	defer closeCfg()

	var c conf.Config
	if err := cfg.Scan(&c); err != nil {
		panic(err)
	}

	newDb(c)

	utils.InitDefaultSnowflakeNode(c.Snowflake.Node)

	grpcServer := server.NewGRPCServer(
		server.WithAddr(c.Server.Grpc.Addr),
		server.WithRedisDsn(c.Data.Redis.Source),
		server.WithRedisPassword(c.Data.Redis.Password),
	)
	app := newBaseService(logger, grpcServer)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
