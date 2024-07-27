package repoiface

import (
	"context"
)

//go:generate mockgen -source=verificationcoderedis.go -destination=verificationcoderedis_mock.go -package=repoiface VerificationCodeRedisRepository
type VerificationCodeRedisRepository interface {
	Insert(ctx context.Context, verificationCodeId, expireTime int64, code string) error
	Get(ctx context.Context, verificationCodeId int64) (string, error)
	Remove(ctx context.Context, verificationCodeId int64) error
}
