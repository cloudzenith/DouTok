package main

import (
	"github.com/cloudzenith/DouTok/backend/gopkgs/launcher"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/server"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
)

func main() {
	launcher.New(
		launcher.WithConfigOptions(
			config.WithSource(file.NewSource("configs/")),
		),
		launcher.WithHttpServer(server.NewHttpServer),
	).Run()
}
