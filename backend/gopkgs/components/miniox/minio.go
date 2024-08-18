package miniox

import (
	"context"
	"fmt"
	"github.com/cloudzenith/DouTok/backend/gopkgs/components"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/minio/minio-go/v6"
	"sync"
)

var (
	globalClientMap = sync.Map{}
	globalConfigMap = make(components.ConfigMap[*Config])
)

func GetConfig() components.ConfigMap[*Config] {
	return globalConfigMap
}

func Init(cm components.ConfigMap[*Config]) (func() error, error) {
	globalConfigMap = cm

	for k, v := range cm {
		client, err := Connect(v)

		if err != nil {
			return nil, err
		}

		globalClientMap.Store(k, client)
	}

	return IsHealth, nil
}

func Connect(c *Config) (*minio.Core, error) {
	c.SetDefault()

	client, err := minio.NewCore(
		fmt.Sprintf("%s:%d", c.Host, c.Port),
		c.AccessKey,
		c.SecretKey,
		c.Secure,
	)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func getKeys(keys ...string) string {
	if len(keys) == 0 {
		return "default"
	}

	return keys[0]
}

func GetClient(ctx context.Context, keys ...string) *minio.Core {
	key := getKeys(keys...)

	value, ok := globalClientMap.Load(key)
	if !ok {
		panic(fmt.Sprintf("minio client %s not found", key))
	}

	return value.(*minio.Core)
}

func IsHealth() (err error) {
	globalClientMap.Range(func(key, value interface{}) bool {
		client := value.(*minio.Core)
		_, err = client.ListBuckets()
		if err != nil {
			log.Error("minio health check failed, client key: %s", key)
			return false
		}

		log.Infof("minio health check success, client key: %s", key)
		return true
	})

	return err
}
