package favoriteappprovider

import (
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/application/favoriteapp"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/service/favoriteservice"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/repositories/favoriterepo"
)

func InitFavoriteApp() *favoriteapp.Application {
	favoriteRepo := favoriterepo.New()
	favoriteService := favoriteservice.New(favoriteRepo)
	favoriteApp := favoriteapp.New(favoriteService)
	return favoriteApp
}
