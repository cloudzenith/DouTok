package data

import (
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/conf"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/pkg/db"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(NewData, NewUserRepo)

type Data struct {
	db  *db.DBClient
	rdb *redis.Client
}

func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	// close default transaction
	dbClient, err := db.NewDBClient(c, &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return nil, nil, err
	}
	return &Data{
		db:  dbClient,
		rdb: nil,
	}, cleanup, nil
}
