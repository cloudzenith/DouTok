//go:build wireinject
// +build wireinject

//go:generate wire

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/conf"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/data"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/server"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/service"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/pkg/utils"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos service.
func wireApp(*conf.Common, *conf.ThirdParty, *conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(utils.ProviderSet, data.ProviderSet, domain.ProviderSet, service.ProviderSet, server.ProviderSet, newApp))
}
