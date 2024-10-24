package videoappproviders

import (
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/applications/videoapp"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/server/commonprovider"
	"github.com/google/wire"
)

var VideoAppProviderSet = wire.NewSet(
	videoapp.New,
	commonprovider.BaseAdapterProvider,
	commonprovider.CoreAdapterProvider,
	commonprovider.VideoServiceProvider,
)
