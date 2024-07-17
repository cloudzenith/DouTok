package repoiface

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/dal/models"
)

type AccountRepository interface {
	Create(ctx context.Context, account *models.Account) error
	ModifyPassword(ctx context.Context, account *models.Account) error
	GetById(ctx context.Context, id int64) (*models.Account, error)
	GetByMobile(ctx context.Context, mobile string) (*models.Account, error)
	GetByEmail(ctx context.Context, email string) (*models.Account, error)
	IsMobileExist(ctx context.Context, mobile string) (bool, error)
	IsEmailExist(ctx context.Context, email string) (bool, error)
}
