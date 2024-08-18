package userdomain

import (
	"context"
	"fmt"
	"github.com/cloudzenith/DouTok/backend/baseService/api"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/conf"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/data/model"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/entity"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/db"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/thirdparty"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/pkg/auth"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/pkg/utils"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
)

type UserUsecase struct {
	config               *conf.Config
	accountServiceClient *thirdparty.AccountServiceClient
	repo                 UserRepo
	snowflake            *utils.SnowflakeNode
	dbClient             *db.DBClient
	log                  *log.Helper
}

func NewUserUsecase(
	config *conf.Config,
	snowflake *utils.SnowflakeNode,
	client *thirdparty.AccountServiceClient,
	repo UserRepo,
	dbClient *db.DBClient,
	logger log.Logger,
) *UserUsecase {
	return &UserUsecase{
		config:               config,
		accountServiceClient: client,
		repo:                 repo,
		snowflake:            snowflake,
		dbClient:             dbClient,
		log:                  log.NewHelper(logger),
	}
}

func (uc *UserUsecase) Register(ctx context.Context, mobile, email, password string) (int64, error) {
	resp, err := uc.accountServiceClient.Register(ctx, &api.RegisterRequest{
		Mobile:   mobile,
		Email:    email,
		Password: password,
	})
	if err != nil {
		return 0, err
	}
	uc.log.Infof("account id: %d", resp.AccountId)

	user := model.User{
		ID:        uc.snowflake.GetSnowflakeId(),
		Mobile:    mobile,
		Email:     email,
		Name:      uuid.New().String(),
		AccountID: resp.AccountId,
	}
	err = uc.dbClient.ExecTx(ctx, func(ctx context.Context) error {
		return uc.repo.Save(ctx, &user)
	})
	if err != nil {
		return 0, err
	}

	return user.ID, nil
}

func (uc *UserUsecase) Login(ctx context.Context, mobile, email, password string) (string, error) {
	resp, err := uc.accountServiceClient.CheckAccount(ctx, &api.CheckAccountRequest{
		Mobile:   mobile,
		Email:    email,
		Password: password,
	})
	if err != nil {
		return "", err
	}

	user, err := uc.repo.FindByAccountID(ctx, resp.AccountId)
	if err != nil {
		return "", err
	}

	token, err := auth.GenerateToken(user.ID, uc.config.Common)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (uc *UserUsecase) UpdateUserInfo(ctx context.Context, user *entity.User) error {
	usermodel := user.ToUserModel()
	var (
		err error
		row int64
	)
	err = uc.dbClient.ExecTx(ctx, func(ctx context.Context) error {
		row, err = uc.repo.UpdateById(ctx, usermodel)
		return err
	})
	if row == 0 {
		return fmt.Errorf("user not found: %d", user.ID)
	}
	return err
}

func (uc *UserUsecase) GetUserInfo(ctx context.Context, userId int64) (*entity.User, error) {
	user, err := uc.repo.FindByID(ctx, userId)
	if err != nil {
		return nil, err
	}
	return entity.FromUserModel(user), err
}
