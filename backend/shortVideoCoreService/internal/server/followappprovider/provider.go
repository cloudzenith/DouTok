package followappprovider

import (
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/application/followapp"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/service/followservice"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/repositories/followrepo"
)

func InitFollowApp() *followapp.Application {
	followRepo := followrepo.New()
	followService := followservice.New(followRepo)
	followApp := followapp.New(followService)
	return followApp
}
