package commonprovider

import (
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/applications/interface/videoserviceiface"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/adapter/baseadapter"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/adapter/svcoreadapter"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/service/videoservice"
	"github.com/google/wire"
)

var CoreAdapterProvider = wire.NewSet(
	svcoreadapter.New,
)

var BaseAdapterProvider = wire.NewSet(
	baseadapter.New,
)

var VideoServiceProvider = wire.NewSet(
	videoservice.New,
	wire.Bind(new(videoserviceiface.VideoService), new(*videoservice.VideoService)),
)
