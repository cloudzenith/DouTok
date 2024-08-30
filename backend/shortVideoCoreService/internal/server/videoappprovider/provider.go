package videoappprovider

import (
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/application/videoapp"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/conf"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/data/userdata"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/data/videodata"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/videodomain"
	"github.com/go-kratos/kratos/v2/log"
)

func InitVideoApplication(config *conf.Config, logger log.Logger) *videoapp.VideoApplication {
	videoRepo := videodata.NewVideoRepo(logger)
	userRepo := userdata.NewUserRepo(logger)
	videoUsecase := videodomain.NewVideoUseCase(config, userRepo, videoRepo, logger)
	videoApp := videoapp.NewVideoApplication(videoUsecase)
	return videoApp
}
