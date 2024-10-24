package collectionappprovider

import (
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/applications/collectionapp"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/server/commonprovider"
	"github.com/google/wire"
)

var CollectionAppProvider = wire.NewSet(
	collectionapp.New,
	commonprovider.CoreAdapterProvider,
	commonprovider.VideoServiceProvider,
)
