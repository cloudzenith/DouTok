package userdomain

import (
	"context"
	"fmt"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/conf"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/data/userdata"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/entity"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/persistence/model"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/persistence/query"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/utils"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
)

type UserUsecase struct {
	config *conf.Config
	repo   userdata.IUserRepo
	log    *log.Helper
}

func NewUserUsecase(
	repo userdata.IUserRepo,
	logger log.Logger,
) *UserUsecase {
	return &UserUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *UserUsecase) CreateUser(ctx context.Context, mobile, email string, accountId int64) (int64, error) {
	user := model.User{
		ID:        utils.GetSnowflakeId(),
		Mobile:    mobile,
		Email:     email,
		Name:      uuid.New().String(),
		AccountID: accountId,
	}
	err := uc.repo.Save(ctx, query.Q, &user)
	if err != nil {
		return 0, err
	}

	return user.ID, nil
}

func (uc *UserUsecase) UpdateUserInfo(ctx context.Context, user *entity.User) error {
	usermodel := user.ToUserModel()
	row, err := uc.repo.UpdateById(ctx, query.Q, usermodel)
	if err != nil {
		return err
	}
	if row == 0 {
		return fmt.Errorf("user not found: %d", user.ID)
	}
	return err
}

func (uc *UserUsecase) GetUserInfo(ctx context.Context, userId int64) (*entity.User, error) {
	user, err := uc.repo.FindByID(ctx, query.Q, userId)
	if err != nil {
		return nil, err
	}
	return entity.FromUserModel(user), err
}
