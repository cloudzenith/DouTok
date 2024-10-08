package userappproviders

import (
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/applications/userapp"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/server/commonprovider"
	"github.com/google/wire"
)

var UserAppProviderSet = wire.NewSet(
	userapp.New,
	commonprovider.BaseAdapterProvider,
	commonprovider.CoreAdapterProvider,
)
