package etcdx

import (
	"context"
	"fmt"
	"github.com/cloudzenith/DouTok/backend/gopkgs/components"
	"github.com/sagikazarmark/crypt/backend/etcd"
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

func Connect(c *Config) (*etcd.ClientV3, error) {
	c.SetDefault()

	client, err := etcd.NewV3([]string{
		c.GetEntrypoint(),
	})

	if err != nil {
		return nil, err
	}

	return client, nil
}

func GetClient(ctx context.Context, keys ...string) *etcd.ClientV3 {
	storeKey := "default"
	if len(keys) != 0 {
		storeKey = keys[0]
	}

	client, ok := globalClientMap.Load(storeKey)
	if !ok {
		panic(fmt.Sprintf("etcd client %s not found", storeKey))
	}

	return client.(*etcd.ClientV3)
}
