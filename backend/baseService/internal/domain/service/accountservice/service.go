package accountservice

import (
	"context"
	"errors"
	"github.com/TremblingV5/box/dbtx"
	"github.com/cloudzenith/DouTok/backend/baseService/api"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/entity/account"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/repoiface"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/dal/models"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/dal/query"
	"gorm.io/gen/field"
)

type Service struct {
	account repoiface.AccountRepository
}

func New(account repoiface.AccountRepository) *Service {
	return &Service{
		account: account,
	}
}

func (s *Service) checkAccountUnique(ctx context.Context, account *account.Account) error {
	if account.Mobile != "" {
		exist, err := s.account.IsMobileExist(ctx, account.Mobile)
		if err != nil {
			return err
		}

		if exist {
			return errors.New("mobile existed")
		}
	}

	if account.Email != "" {
		exist, err := s.account.IsEmailExist(ctx, account.Email)
		if err != nil {
			return err
		}

		if exist {
			return errors.New("email existed")
		}
	}

	return nil
}

func (s *Service) Create(ctx context.Context, mobile, email, password string) (int64, error) {
	account := account.NewAccount(
		account.WithMobile(mobile),
		account.WithEmail(email),
		account.WithPassword(password),
	)

	if err := s.checkAccountUnique(ctx, account); err != nil {
		return 0, err
	}

	if isValid := account.IsPasswordValid(); !isValid {
		return 0, errors.New("invalid password")
	}

	if err := account.EncryptPassword(); err != nil {
		return 0, err
	}

	account.GenerateId()
	if err := s.account.Create(ctx, account.ToModel()); err != nil {
		return 0, err
	}

	return account.ID, nil
}

func (s *Service) checkPassword(ctx context.Context, getDataFunc func() (*models.Account, error), password string) (*account.Account, error) {
	accountDo, err := getDataFunc()
	if err != nil {
		return nil, err
	}

	account := account.NewAccountWithModel(accountDo)
	if err := account.CheckPassword(password); err != nil {
		return nil, err
	}

	return account, nil
}

func (s *Service) CheckPasswordById(ctx context.Context, id int64, password string) (int64, error) {
	account, err := s.checkPassword(ctx, func() (*models.Account, error) {
		return s.account.GetById(ctx, id)
	}, password)

	if err != nil {
		return 0, err
	}

	return account.ID, err
}

func (s *Service) CheckPasswordByMobile(ctx context.Context, mobile, password string) (int64, error) {
	account, err := s.checkPassword(ctx, func() (*models.Account, error) {
		return s.account.GetByMobile(ctx, mobile)
	}, password)

	if err != nil {
		return 0, err
	}

	return account.ID, err
}

func (s *Service) CheckPasswordByEmail(ctx context.Context, email, password string) (int64, error) {
	account, err := s.checkPassword(ctx, func() (*models.Account, error) {
		return s.account.GetByEmail(ctx, email)
	}, password)

	if err != nil {
		return 0, err
	}

	return account.ID, err
}

func (s *Service) ModifyPassword(ctx context.Context, id int64, oldPassword, newPassword string) error {
	account, err := s.checkPassword(ctx, func() (*models.Account, error) {
		return s.account.GetById(ctx, id)
	}, oldPassword)
	if err != nil {
		return err
	}

	if err := account.ModifyPassword(newPassword); err != nil {
		return err
	}

	if err := s.account.ModifyPassword(ctx, account.ToModel()); err != nil {
		return err
	}

	return nil
}

func (s *Service) Unbind(ctx context.Context, id int64, voucherType api.VoucherType) (err error) {
	ctx, persist := dbtx.WithTXPersist(ctx)
	defer func() {
		persist(err)
	}()

	if id == 0 {
		return errors.New("账户id不能为空")
	}

	var column field.Expr
	switch voucherType {
	case api.VoucherType_VOUCHER_EMAIL:
		column = query.Q.Account.Email
	case api.VoucherType_VOUCHER_PHONE:
		column = query.Q.Account.Mobile
	default:
		return errors.New("不支持的类型")
	}

	return s.account.ClearColumn(ctx, column)
}

func (s *Service) Bind(ctx context.Context, id int64, voucherType api.VoucherType, voucher string) (err error) {
	ctx, persist := dbtx.WithTXPersist(ctx)
	defer func() {
		persist(err)
	}()

	if id == 0 {
		return errors.New("账户id不能为空")
	}

	var column field.Expr
	switch voucherType {
	case api.VoucherType_VOUCHER_EMAIL:
		column = query.Q.Account.Email
	case api.VoucherType_VOUCHER_PHONE:
		column = query.Q.Account.Mobile
	default:
		return errors.New("不支持的类型")
	}

	return s.account.UpdateColumn(ctx, column, voucher)
}
