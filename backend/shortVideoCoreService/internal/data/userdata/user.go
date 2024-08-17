package userdata

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/data/model"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/db"
	"github.com/go-kratos/kratos/v2/log"
)

type UserRepo struct {
	dbClient *db.DBClient
	log      *log.Helper
}

// NewUserRepo .
func NewUserRepo(dbClient *db.DBClient, logger log.Logger) *UserRepo {
	return &UserRepo{
		dbClient: dbClient,
		log:      log.NewHelper(logger),
	}
}

func (r *UserRepo) Save(ctx context.Context, u *model.User) error {
	result := r.dbClient.DB(ctx).Create(u)
	return result.Error
}

func (r *UserRepo) UpdateById(ctx context.Context, u *model.User) (int64, error) {
	result := r.dbClient.DB(ctx).Where(&model.User{ID: u.ID}).Updates(u)
	return result.RowsAffected, result.Error
}

func (r *UserRepo) FindByID(ctx context.Context, userID int64) (*model.User, error) {
	user := &model.User{}
	result := r.dbClient.DB(ctx).First(user, userID)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (r *UserRepo) FindByAccountID(ctx context.Context, accountID int64) (*model.User, error) {
	user := &model.User{}
	result := r.dbClient.DB(ctx).Where(&model.User{AccountID: accountID}).First(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
