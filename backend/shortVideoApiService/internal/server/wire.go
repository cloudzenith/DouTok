//go:build wireinject
// +build wireinject

package server

import (
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/applications/collectionapp"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/applications/commentapp"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/applications/favoriteapp"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/applications/fileapp"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/applications/followapp"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/applications/userapp"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/applications/videoapp"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/server/collectionappprovider"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/server/commentappprovider"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/server/favoriteappprovider"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/server/fileappproviders"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/server/followappprovider"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/server/userappproviders"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/server/videoappproviders"
	"github.com/google/wire"
)

func initUserApp() *userapp.Application {
	wire.Build(userappproviders.UserAppProviderSet)
	return &userapp.Application{}
}

func initVideoApp() *videoapp.Application {
	wire.Build(videoappproviders.VideoAppProviderSet)
	return &videoapp.Application{}
}

func initFileApp() *fileapp.Application {
	wire.Build(fileappproviders.FileAppProviderSet)
	return &fileapp.Application{}
}

func initCollectionApp() *collectionapp.Application {
	wire.Build(collectionappprovider.CollectionAppProvider)
	return &collectionapp.Application{}
}

func initCommentApp() *commentapp.Application {
	wire.Build(commentappprovider.CommentAppProvider)
	return &commentapp.Application{}
}

func initFavoriteApp() *favoriteapp.Application {
	wire.Build(favoriteappprovider.FavoriteAppProvider)
	return &favoriteapp.Application{}
}

func initFollowApp() *followapp.Application {
	wire.Build(followappprovider.FollowAppProvider)
	return &followapp.Application{}
}
