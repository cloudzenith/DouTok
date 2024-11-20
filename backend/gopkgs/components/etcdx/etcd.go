package etcdx

import (
	"context"
	"fmt"
	"github.com/cloudzenith/DouTok/backend/gopkgs/components"
	"github.com/go-kratos/kratos/v2/log"
	clientv3 "go.etcd.io/etcd/client/v3"
	"sync"
	"time"
)

const healthCheckTimeout = 3 * time.Second

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

func Connect(c *Config) (*clientv3.Client, error) {
	c.SetDefault()

	client, err := clientv3.New(clientv3.Config{
		Endpoints: []string{
			c.GetEntrypoint(),
		},
	})

	if err != nil {
		return nil, err
	}

	return client, nil
}

func GetClient(ctx context.Context, keys ...string) *clientv3.Client {
	storeKey := "default"
	if len(keys) != 0 {
		storeKey = keys[0]
	}

	client, ok := globalClientMap.Load(storeKey)
	if !ok {
		panic(fmt.Sprintf("etcd client %s not found", storeKey))
	}

	return client.(*clientv3.Client)
}

func IsHealth() (err error) {
	globalClientMap.Range(func(key, value interface{}) bool {
		client := value.(*clientv3.Client)
		timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), healthCheckTimeout)
		defer cancelFunc()
		_, err = client.Get(timeoutCtx, "health")
		if err != nil {
			log.Errorf("etcd health check failed, client key: %s", key)
			return false
		}

		log.Infof("etcd health check success, client key: %s", key)
		return true
	})

	return err
}
