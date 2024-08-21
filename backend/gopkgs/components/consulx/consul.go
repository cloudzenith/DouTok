package consulx

import (
	"context"
	"fmt"
	"github.com/cloudzenith/DouTok/backend/gopkgs/components"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/hashicorp/consul/api"
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

func Connect(c *Config) (*api.Client, error) {
	c.SetDefault()

	cfg := api.DefaultConfig()
	cfg.Address = c.Address

	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		return nil, err
	}

	return client, nil
}

func GetClient(ctx context.Context, keys ...string) *api.Client {
	storeKey := "default"
	if len(keys) != 0 {
		storeKey = keys[0]
	}

	client, ok := globalClientMap.Load(storeKey)
	if !ok {
		panic(fmt.Sprintf("consul client %s not found", storeKey))
	}

	return client.(*api.Client)
}

func IsHealth() (err error) {
	globalClientMap.Range(func(key, value interface{}) bool {
		client := value.(*api.Client)
		_, _, err = client.Health().State("any", nil)
		if err != nil {
			log.Errorf("consul health check failed, client key: %s", key)
			return false
		}

		log.Infof("consul health check success, client key: %s", key)
		return true
	})

	return err
}
