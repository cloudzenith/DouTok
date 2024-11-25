package main

import (
	"github.com/cloudzenith/DouTok/backend/gopkgs/launcher"
	"github.com/cloudzenith/DouTok/backend/manageApiService/internal/conf"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
)

func main() {
	c := &conf.Config{}

	launcher.New(
		launcher.WithConfigValue(c),
		launcher.WithConfigOptions(
			config.WithSource(file.NewSource("configs/")),
		),
	).Run()
}
