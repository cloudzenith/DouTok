//go:build wireinject
// +build wireinject

package server

import (
	"github.com/cloudzenith/DouTok/backend/baseService/internal/applications/accountapp"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/applications/authapp"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/applications/fileapp"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/applications/postapp"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/innerservice/filerepohelper"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/server/accountproviders"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/server/authappproviders"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/server/fileappproviders"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/server/postappproviders"
	"github.com/google/wire"
)

func initAccountApplication() *accountapp.AccountApplication {
	wire.Build(accountproviders.AccountAppProviderSet)
	return nil
}

func initAuthApplication() *authapp.AuthApplication {
	wire.Build(authappproviders.AuthAppProviderSet)
	return nil
}

func initPostApplication() *postapp.PostApplication {
	wire.Build(postappproviders.PostAppProviderSet)
	return nil
}

func initFileApplication(fileTableShardingConfig filerepohelper.FileTableShardingConfig) *fileapp.FileApplication {
	wire.Build(fileappproviders.FileAppProviderSet)
	return nil
}
