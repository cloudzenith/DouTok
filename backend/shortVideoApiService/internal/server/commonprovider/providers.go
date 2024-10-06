package commonprovider

import (
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/adapter/baseadapter"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/adapter/svcoreadapter"
	"github.com/google/wire"
)

var CoreAdapterProvider = wire.NewSet(
	svcoreadapter.New,
)

var BaseAdapterProvider = wire.NewSet(
	baseadapter.New,
)
