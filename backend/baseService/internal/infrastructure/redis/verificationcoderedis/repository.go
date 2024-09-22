package verificationcoderedis

import (
	"context"
	"fmt"
	"github.com/cloudzenith/DouTok/backend/gopkgs/components/redisx"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisRepository struct {
	db *redis.Client
}

func New() *RedisRepository {
	return &RedisRepository{db: redisx.GetClient(context.Background())}
}

func (r *RedisRepository) formatVerificationCodeKey(verificationCodeId int64) string {
	return fmt.Sprintf("VERIFICATION_CODE_ID_%d", verificationCodeId)
}

func (r *RedisRepository) Insert(ctx context.Context, verificationCodeId, expireTime int64, code string) error {
	key := r.formatVerificationCodeKey(verificationCodeId)
	_, err := r.db.Set(ctx, key, code, time.Duration(expireTime)*time.Millisecond).Result()
	log.Context(ctx).Infof("insert verification key: %s, code: %s, seconds: %v", key, code, time.Duration(expireTime)*time.Second)
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisRepository) Get(ctx context.Context, verificationCodeId int64) (string, error) {
	code, err := r.db.Get(ctx, r.formatVerificationCodeKey(verificationCodeId)).Result()
	if err != nil {
		return "", err
	}

	return code, nil
}

func (r *RedisRepository) Remove(ctx context.Context, verificationCodeId int64) error {
	_, err := r.db.Del(ctx, r.formatVerificationCodeKey(verificationCodeId)).Result()
	if err != nil {
		return err
	}

	return nil
}
