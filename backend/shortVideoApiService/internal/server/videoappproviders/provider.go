package videoappproviders

import (
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/applications/videoapp"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/adapter/baseadapter"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/adapter/svcoreadapter"
	"github.com/google/wire"
)

var BaseAdapterProvider = wire.NewSet(
	baseadapter.New,
)

var CoreAdapterProvider = wire.NewSet(
	svcoreadapter.New,
)

var VideoAppProviderSet = wire.NewSet(
	videoapp.New,
	BaseAdapterProvider,
	CoreAdapterProvider,
)
