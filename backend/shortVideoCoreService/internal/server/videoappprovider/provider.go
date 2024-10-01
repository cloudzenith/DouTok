package videoappprovider

import (
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/application/videoapp"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/conf"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/data/userdata"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/data/videodata"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/service/videodomain"
)

func InitVideoApplication(config *conf.Config) *videoapp.VideoApplication {
	videoRepo := videodata.NewVideoRepo()
	userRepo := userdata.NewUserRepo()
	videoUsecase := videodomain.NewVideoUseCase(config, userRepo, videoRepo)
	videoApp := videoapp.NewVideoApplication(videoUsecase)
	return videoApp
}
