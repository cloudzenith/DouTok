package accountrepo

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/dal/query"
	"gorm.io/gen"

	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/dal/models"
)

type PersistRepository struct{}

func New() *PersistRepository {
	return &PersistRepository{}
}

func (r *PersistRepository) getFirst(ctx context.Context, conditions ...gen.Condition) (*models.Account, error) {
	conditions = append(conditions, query.Q.Account.IsDeleted.Is(false))
	account, err := query.Q.WithContext(ctx).Account.Where(
		conditions...,
	).First()
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (r *PersistRepository) Create(ctx context.Context, account *models.Account) error {
	return query.Q.WithContext(ctx).Account.Create(account)
}

func (r *PersistRepository) ModifyPassword(ctx context.Context, account *models.Account) error {
	_, err := query.Q.WithContext(ctx).Account.UpdateColumns(account)
	return err
}

func (r *PersistRepository) GetById(ctx context.Context, id int64) (*models.Account, error) {
	return r.getFirst(ctx, query.Q.Account.ID.Eq(id))
}

func (r *PersistRepository) GetByMobile(ctx context.Context, mobile string) (*models.Account, error) {
	return r.getFirst(ctx, query.Q.Account.Mobile.Eq(mobile))
}

func (r *PersistRepository) GetByEmail(ctx context.Context, email string) (*models.Account, error) {
	return r.getFirst(ctx, query.Q.Account.Email.Eq(email))
}

func (r *PersistRepository) count(ctx context.Context, conditions ...gen.Condition) (int64, error) {
	conditions = append(conditions, query.Q.Account.IsDeleted.Is(false))
	return query.Q.WithContext(ctx).Account.Where(
		conditions...,
	).Count()
}

func (r *PersistRepository) isExist(ctx context.Context, conditions ...gen.Condition) (bool, error) {
	count, err := r.count(ctx, conditions...)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *PersistRepository) IsMobileExist(ctx context.Context, mobile string) (bool, error) {
	return r.isExist(ctx, query.Q.Account.Mobile.Eq(mobile))
}

func (r *PersistRepository) IsEmailExist(ctx context.Context, email string) (bool, error) {
	return r.isExist(ctx, query.Q.Account.Email.Eq(email))
}
