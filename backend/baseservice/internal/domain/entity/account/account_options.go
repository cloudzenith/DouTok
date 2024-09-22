package account

type AccountOption func(*Account)

func WithID(id int64) AccountOption {
	return func(a *Account) {
		a.ID = id
	}
}

func WithMobile(mobile string) AccountOption {
	return func(a *Account) {
		a.Mobile = mobile
	}
}

func WithEmail(email string) AccountOption {
	return func(a *Account) {
		a.Email = email
	}
}

func WithPassword(password string) AccountOption {
	return func(a *Account) {
		a.Password = password
	}
}
