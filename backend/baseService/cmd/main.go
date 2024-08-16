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
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
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

func newDb(cfg conf.Config) *gorm.DB {
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
	return db
}

func newMinioCore(cfg conf.Config) *minio.Core {
	core, err := minio.NewCore(
		cfg.Data.Minio.Endpoint,
		&minio.Options{
			Creds: credentials.NewStaticV4(
				cfg.Data.Minio.AccessKey,
				cfg.Data.Minio.SecretKey,
				"",
			),
		},
	)
	if err != nil {
		panic(err)
	}

	return core
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

	db := newDb(c)
	core := newMinioCore(c)
	utils.InitDefaultSnowflakeNode(c.Snowflake.Node)

	grpcServer := server.NewGRPCServer(
		server.WithAddr(c.Server.Grpc.Addr),
		server.WithRedisDsn(c.Data.Redis.Source),
		server.WithRedisPassword(c.Data.Redis.Password),
		server.WithDB(db),
		server.WithMinioCore(core),
		server.WithFileTableShardingConfig(c.Data),
		server.WithDBShardingTablesConfig(c.Data.DbShardingTables),
	)
	app := newBaseService(logger, grpcServer)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
