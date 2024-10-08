package followappprovider

import (
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/applications/followapp"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/server/commonprovider"
	"github.com/google/wire"
)

var FollowAppProvider = wire.NewSet(
	followapp.New,
	commonprovider.CoreAdapterProvider,
)
