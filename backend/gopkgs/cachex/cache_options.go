package cachex

import (
	"github.com/redis/go-redis/v9"
	"time"
)

type CacheOption[T any] func(cache *Cache[T])

func WithUseBarrier[T any](useBarrier bool) CacheOption[T] {
	return func(cache *Cache[T]) {
		cache.useBarrier = useBarrier
	}
}

func WithUseLocal[T any](useLocal bool) CacheOption[T] {
	return func(cache *Cache[T]) {
		cache.useLocal = useLocal
	}
}

func WithUseFallback[T any](useFallback bool) CacheOption[T] {
	return func(cache *Cache[T]) {
		cache.useFallback = useFallback
	}
}

func WithLocalTTL[T any](localTTL *time.Duration) CacheOption[T] {
	return func(cache *Cache[T]) {
		cache.localTTL = localTTL
	}
}

func WithRedisTTL[T any](redisTTL *time.Duration) CacheOption[T] {
	return func(cache *Cache[T]) {
		cache.redisTTL = redisTTL
	}
}

func WithRedisClient[T any](client *redis.Client) CacheOption[T] {
	return func(cache *Cache[T]) {
		cache.client = client
	}
}
