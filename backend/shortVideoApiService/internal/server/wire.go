//go:build wireinject
// +build wireinject

package server

import (
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/applications/userapp"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/server/userappproviders"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/server/videoappproviders"
	"github.com/google/wire"
)

func initUserApp() *userapp.Application {
	wire.Build(userappproviders.UserAppProviderSet)
	return &userapp.Application{}
}

func initVideoApp() *userapp.Application {
	wire.Build(videoappproviders.VideoAppProviderSet)
	return &userapp.Application{}
}
