package userdata

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/persistence/model"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/persistence/query"
)

type UserRepo struct {
}

// NewUserRepo .
func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (r *UserRepo) Save(ctx context.Context, tx *query.Query, u *model.User) error {
	return tx.WithContext(ctx).User.Create(u)
}

func (r *UserRepo) UpdateById(ctx context.Context, tx *query.Query, u *model.User) (int64, error) {
	ret, err := tx.User.WithContext(ctx).Where(tx.User.ID.Eq(u.ID)).Updates(u)
	if err != nil {
		return 0, err
	}
	return ret.RowsAffected, nil
}

func (r *UserRepo) FindByID(ctx context.Context, tx *query.Query, userID int64) (*model.User, error) {
	user, err := tx.User.WithContext(ctx).Where(tx.User.ID.Eq(userID)).First()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepo) FindByAccountID(ctx context.Context, tx *query.Query, accountID int64) (*model.User, error) {
	user, err := tx.User.WithContext(ctx).Where(tx.User.AccountID.Eq(accountID)).First()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepo) FindByIds(ctx context.Context, tx *query.Query, ids []int64) ([]*model.User, error) {
	users, err := tx.User.WithContext(ctx).Where(tx.User.ID.In(ids...)).Find()
	if err != nil {
		return nil, err
	}
	return users, nil
}
