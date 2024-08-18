package miniox

import (
	"context"
	"fmt"
	"github.com/TremblingV5/box/components"
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

func Init(cm components.ConfigMap[*Config]) error {
	globalConfigMap = cm

	for k, v := range cm {
		client, err := Connect(v)

		if err != nil {
			return err
		}

		globalClientMap.Store(k, client)
	}

	return nil
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
