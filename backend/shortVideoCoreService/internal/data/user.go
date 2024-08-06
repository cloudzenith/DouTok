package data

import (
	"context"
	"fmt"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/data/model"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/do"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/pkg/db"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type userRepo struct {
	storage *Data
	log     *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) domain.UserRepo {
	return &userRepo{
		storage: data,
		log:     log.NewHelper(logger),
	}
}

func (r *userRepo) Save(ctx context.Context, tx *gorm.DB, u *do.User) error {
	if tx == nil {
		return fmt.Errorf("transaction not provided")
	}

	user := u.ToUserModel()
	tx = tx.Create(&user)
	return tx.Error
}

func (r *userRepo) Update(ctx context.Context, tx *gorm.DB, u *do.User) error {
	if tx == nil {
		return fmt.Errorf("transaction not provided")
	}

	user, err := r.FindByID(ctx, tx, u.ID)
	if err != nil {
		return err
	}
	newUser := u.ToUserModel()
	tx = tx.Model(&user).Updates(newUser)
	return tx.Error
}

func (r *userRepo) FindByID(ctx context.Context, tx *gorm.DB, userID int64) (*do.User, error) {
	if tx == nil {
		return nil, fmt.Errorf("transaction not provided")
	}

	user := model.User{}
	tx = tx.First(&user, userID)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return do.FromUserModel(user), nil
}

func (r *userRepo) StartTransaction(ctx context.Context) (*gorm.DB, *db.TransactionMaker, error) {
	if r.storage.db == nil {
		return nil, nil, fmt.Errorf("database not provided")
	}
	return r.storage.db.StartTransaction()
}
