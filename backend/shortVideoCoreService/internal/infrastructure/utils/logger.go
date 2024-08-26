package utils

import (
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"os"
)

func InitStdLogger(config *conf.App) log.Logger {
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.name", config.Name,
		"service.version", config.Version,
		"service.node", config.Node,
	)
	return logger
}
