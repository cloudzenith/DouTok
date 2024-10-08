package collectionappprovider

import (
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/application/collectionapp"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/service/collectionservice"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/repositories/collectionrepo"
)

func InitCollectionApplication() *collectionapp.Application {
	collectionRepo := collectionrepo.New()
	collectionService := collectionservice.New(collectionRepo)
	collectionApp := collectionapp.New(collectionService)
	return collectionApp
}
