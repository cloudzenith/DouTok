package authappproviders

import (
	"github.com/cloudzenith/DouTok/backend/baseService/internal/applications/authapp"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/applications/interface/authserviceiface"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/repoiface"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/service/authservice"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/redis/verificationcoderedis"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
)

type RedisDsn string
type RedisPassword string

func NewRedis(dsn RedisDsn, password RedisPassword) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     string(dsn),
		Password: string(password),
	})
}

var VerificationCodeRedisRepositoryProviders = wire.NewSet(
	verificationcoderedis.New,
	NewRedis,
	wire.Bind(new(repoiface.VerificationCodeRedisRepository), new(*verificationcoderedis.RedisRepository)),
)

var AuthServiceProviders = wire.NewSet(
	authservice.New,
	wire.Bind(new(authserviceiface.AuthService), new(*authservice.AuthService)),
)

var AuthAppProviderSet = wire.NewSet(
	authapp.New,
	VerificationCodeRedisRepositoryProviders,
	AuthServiceProviders,
)
