package cachex

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/jellydator/ttlcache/v3"
	"time"

	"github.com/cloudzenith/DouTok/backend/gopkgs/gofer"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
)

const (
	NotFoundBarrier = "{404}"
	defaultTTL      = 5 * time.Second
)

var (
	ErrNotFoundBarrier = errors.New("err not found barrier")
	DefaultTTL         = defaultTTL
)

type Cache[T any] struct {
	useBarrier  bool
	useLocal    bool
	useFallback bool
	localTTL    *time.Duration
	redisTTL    *time.Duration
	client      *redis.Client
	localCache  *ttlcache.Cache[string, []byte]
}

func NewCache[T any](options ...CacheOption[T]) *Cache[T] {
	cache := &Cache[T]{}

	for _, option := range options {
		option(cache)
	}

	if cache.client == nil {
		panic("not assign redis client to cache")
	}

	if cache.localTTL == nil {
		cache.localTTL = &DefaultTTL
	}

	if cache.redisTTL == nil {
		cache.redisTTL = &DefaultTTL
	}

	if cache.useLocal {
		cache.localCache = ttlcache.New[string, []byte]()
	}

	return cache
}

func (c *Cache[T]) fetchSet(ctx context.Context, key string, fetch func(ctx context.Context) (T, error)) (t T, err error) {
	value, err := fetch(ctx)
	if err != nil {
		if c.useBarrier {
			_ = c.client.Set(ctx, key, NotFoundBarrier, *c.redisTTL).Err()
			return t, ErrNotFoundBarrier
		}
		return t, err
	}

	val, err := json.Marshal(value)
	if err != nil {
		return t, err
	}

	ttl := c.redisTTL
	err = c.client.Set(ctx, key, val, *ttl).Err()
	if err != nil {
		log.Context(ctx).Warnf("fetchSet write redis failed: %v", err)
	}

	if c.useLocal {
		if *c.localTTL > 0 {
			ttl = c.localTTL
		}

		c.localCache.Set(key, val, *ttl)
	}

	return value, nil
}

func (c *Cache[T]) getCache(ctx context.Context, key string) ([]byte, error) {
	if c.useLocal {
		item := c.localCache.Get(key)
		if item != nil {
			return item.Value(), nil
		}
	}

	result, err := c.client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	if result == NotFoundBarrier {
		return nil, ErrNotFoundBarrier
	}

	return []byte(result), nil
}

func (c *Cache[T]) Fetch(
	ctx context.Context,
	key string,
	fetch func(ctx context.Context) (T, error),
) (t T, err error) {
	value, err, _ := gofer.SingleFlightDo(key, func() (any, error) {
		if c.useFallback {
			val, err := c.fetchSet(ctx, key, fetch)
			if err == nil {
				return val, nil
			}

			return c.getCache(ctx, key)
		}

		val, err := c.getCache(ctx, key)
		if err == nil {
			return val, nil
		}

		if errors.Is(err, ErrNotFoundBarrier) {
			return nil, err
		}

		return c.fetchSet(ctx, key, fetch)
	})

	gofer.SingleFlightForget(key)
	if err != nil {
		return t, err
	}

	err = json.Unmarshal(value.([]byte), t)
	if err != nil {
		return t, err
	}

	return t, nil
}
