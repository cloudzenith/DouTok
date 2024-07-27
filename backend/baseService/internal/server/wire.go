//go:build wireinject
// +build wireinject

package server

import (
	"github.com/cloudzenith/DouTok/backend/baseService/internal/applications/accountapp"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/applications/authapp"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/server/accountproviders"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/server/authappproviders"
	"github.com/google/wire"
)

func initAccountApplication() *accountapp.AccountApplication {
	wire.Build(accountproviders.AccountAppProviderSet)
	return nil
}

func initAuthApplication(dsn authappproviders.RedisDsn, password authappproviders.RedisPassword) *authapp.AuthApplication {
	wire.Build(authappproviders.AuthAppProviderSet)
	return nil
}
