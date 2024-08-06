package domain

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/baseService/api"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/do"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/pkg/db"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/pkg/utils"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

// UserRepo repo 层是抽象，这里与 data 层的实现有关，但是需要做到事务跨多个sql，所以需要在 repo 层定义事务方法，可以用
// context 传递事务，也可以用 gorm 的事务方法？抽象一个 session 的概念，用于事务管理
type UserRepo interface {
	Save(ctx context.Context, tx *gorm.DB, u *do.User) error
	Update(ctx context.Context, tx *gorm.DB, u *do.User) error
	FindByID(ctx context.Context, tx *gorm.DB, id int64) (*do.User, error)
	StartTransaction(ctx context.Context) (*gorm.DB, *db.TransactionMaker, error)
}

type UserUsecase struct {
	baseServiceClient api.AccountServiceClient
	repo              UserRepo
	snowflake         *utils.SnowflakeNode
	log               *log.Helper
}

func NewUserUsecase(
	snowflake *utils.SnowflakeNode,
	client api.AccountServiceClient,
	repo UserRepo,
	logger log.Logger) *UserUsecase {
	return &UserUsecase{
		baseServiceClient: client,
		repo:              repo,
		snowflake:         snowflake,
		log:               log.NewHelper(logger),
	}
}

func (uc *UserUsecase) Register(ctx context.Context, u *do.User) (*do.User, error) {
	tx, maker, err := uc.repo.StartTransaction(ctx)
	if err != nil {
		return nil, err
	}
	defer maker.Close(&err)

	//通过 baseService 创建 account
	resp, err := uc.baseServiceClient.Register(ctx, &api.RegisterRequest{
		Mobile:   u.Mobile,
		Email:    u.Email,
		Password: u.Password,
	})
	if err != nil {
		return nil, err
	}

	u.AccountID = resp.AccountId
	u.ID = uc.snowflake.GetSnowflakeId()

	err = uc.repo.Save(ctx, tx, u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (uc *UserUsecase) Login(ctx context.Context, u *do.User) (string, error) {
	// TODO: 生成 JWT token
	_, err := uc.baseServiceClient.CheckAccount(ctx, &api.CheckAccountRequest{
		Mobile:   u.Mobile,
		Email:    u.Email,
		Password: u.Password,
	})
	if err != nil {
		return "", err
	}
	return "token", nil
}

func (uc *UserUsecase) UpdateUserInfo(ctx context.Context, u *do.User) (*do.User, error) {
	tx, maker, err := uc.repo.StartTransaction(ctx)
	if err != nil {
		return nil, err
	}
	defer maker.Close(&err)

	err = uc.repo.Update(ctx, tx, u)
	if err != nil {
		return nil, err
	}
	return nil, err
}

func (uc *UserUsecase) GetUserInfo(ctx context.Context, id int64) (*do.User, error) {
	tx, maker, err := uc.repo.StartTransaction(ctx)
	if err != nil {
		return nil, err
	}
	defer maker.Close(&err)

	user, err := uc.repo.FindByID(ctx, tx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
