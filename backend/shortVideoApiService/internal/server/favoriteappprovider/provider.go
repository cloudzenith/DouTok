package favoriteappprovider

import (
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/applications/favoriteapp"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/server/commonprovider"
	"github.com/google/wire"
)

var FavoriteAppProvider = wire.NewSet(
	favoriteapp.New,
	commonprovider.CoreAdapterProvider,
	commonprovider.VideoServiceProvider,
)
