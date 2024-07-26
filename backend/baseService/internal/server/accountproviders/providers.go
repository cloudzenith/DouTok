package accountproviders

import (
	"github.com/cloudzenith/DouTok/backend/baseService/internal/applications/accountapp"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/applications/interface/accountserviceiface"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/repoiface"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/service/accountservice"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/repositories/accountrepo"
	"github.com/google/wire"
)

var AccountRepoProviders = wire.NewSet(
	accountrepo.New,
	wire.Bind(new(repoiface.AccountRepository), new(*accountrepo.PersistRepository)),
)

var AccountServiceProviders = wire.NewSet(
	accountservice.New,
	wire.Bind(new(accountserviceiface.AccountService), new(*accountservice.Service)),
)

var AccountAppProviderSet = wire.NewSet(
	accountapp.New,
	AccountRepoProviders,
	AccountServiceProviders,
)
