package account

import (
	"errors"

	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/constants"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/dal/models"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/utils"
)

const (
	ErrInvalidPassword = "密码由大小写字母、数字、符号组成，且至少需要8位"
)

type Account struct {
	ID       int64
	Mobile   string
	Email    string
	Password string
	Salt     string
}

func NewAccount(options ...AccountOption) *Account {
	account := &Account{}
	for _, option := range options {
		option(account)
	}

	return account
}

func NewAccountWithModel(model *models.Account) *Account {
	return &Account{
		ID:       model.ID,
		Mobile:   model.Mobile,
		Email:    model.Email,
		Password: model.Password,
		Salt:     model.Salt,
	}
}

func (a *Account) ToModel() *models.Account {
	return &models.Account{
		ID:       a.ID,
		Mobile:   a.Mobile,
		Email:    a.Email,
		Password: a.Password,
		Salt:     a.Salt,
	}
}

func (a *Account) IsPasswordValid(patterns ...string) bool {
	// check with the given pattern
	if len(patterns) > 0 {
		pattern := patterns[0]
		return utils.IsValidWithRegex(pattern, a.Password)
	}

	// check with the default pattern
	return utils.IsValidWithRegex(constants.AccountPasswordPattern, a.Password)
}

func (a *Account) ModifyPassword(password string) error {
	a.Password = password

	isValid := a.IsPasswordValid()
	if !isValid {
		return errors.New(ErrInvalidPassword)
	}

	if err := a.EncryptPassword(); err != nil {
		return err
	}

	return nil
}

func (a *Account) EncryptPassword() error {
	if err := a.generateSalt(); err != nil {
		return err
	}

	a.Password = utils.GenerateMd5WithSalt(a.Password, a.Salt)
	return nil
}

func (a *Account) generateSalt() error {
	salt, err := utils.GetPasswordSalt()
	if err != nil {
		return err
	}

	a.Salt = salt
	return nil
}

func (a *Account) CheckPassword(password string) error {
	passwordMd5 := utils.GenerateMd5WithSalt(password, a.Salt)
	if passwordMd5 != a.Password {
		return errors.New("wrong password")
	}

	return nil
}

func (a *Account) GenerateId() {
	a.ID = utils.GetSnowflakeId()
}
