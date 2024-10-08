package commentappprovider

import (
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/applications/commentapp"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/server/commonprovider"
	"github.com/google/wire"
)

var CommentAppProvider = wire.NewSet(
	commentapp.New,
	commonprovider.CoreAdapterProvider,
)
