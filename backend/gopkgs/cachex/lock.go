package cachex

import (
	"context"
	"fmt"
	"github.com/bsm/redislock"
	"github.com/redis/go-redis/v9"
	"time"
)

type LockHandle struct {
	locker      *redislock.Client
	client      *redis.Client
	lockPattern string
	newLockTTL  time.Duration
	renewTTL    time.Duration
}

func New(client *redis.Client, lockPattern string, newLockTTL, renewTTL time.Duration) *LockHandle {
	return &LockHandle{
		locker:      redislock.New(client),
		client:      client,
		lockPattern: lockPattern,
		newLockTTL:  newLockTTL,
		renewTTL:    renewTTL,
	}
}

func (l *LockHandle) Lock(ctx context.Context, keywords ...any) (*redislock.Lock, error) {
	lockKey := fmt.Sprintf(l.lockPattern, keywords...)
	return l.locker.Obtain(ctx, lockKey, l.newLockTTL, nil)
}
