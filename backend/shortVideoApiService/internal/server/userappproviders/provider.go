package userappproviders

import (
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/applications/userapp"
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

var UserAppProviderSet = wire.NewSet(
	userapp.New,
	BaseAdapterProvider,
	CoreAdapterProvider,
)
