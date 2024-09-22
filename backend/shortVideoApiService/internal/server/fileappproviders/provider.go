package fileappproviders

import (
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/applications/fileapp"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/adapter/baseadapter"
	"github.com/google/wire"
)

var BaseAdapterProvider = wire.NewSet(
	baseadapter.New,
)

var FileAppProviderSet = wire.NewSet(
	fileapp.New,
	BaseAdapterProvider,
)
