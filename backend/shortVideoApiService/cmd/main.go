package main

import (
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/server"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
)

var logger log.Logger

func newBaseService(logger log.Logger, gs *http.Server) *kratos.App {
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

func main() {
	//cfg, closeCfg := loadConfig()
	//defer closeCfg()

	ginServer := server.NewGinServer()
	app := newBaseService(logger, ginServer)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
