package userappprovider

import (
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/application/userapp"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/conf"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/data/userdata"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/userdomain"
	"github.com/go-kratos/kratos/v2/log"
)

func InitUserApplication(config *conf.Config, logger log.Logger) *userapp.UserApplication {
	userRepo := userdata.NewUserRepo(logger)
	userUsecase := userdomain.NewUserUsecase(userRepo, logger)
	userApp := userapp.NewUserApplication(userUsecase)
	return userApp
}
