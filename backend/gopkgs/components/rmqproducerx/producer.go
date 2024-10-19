package rmqproducerx

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/cloudzenith/DouTok/backend/gopkgs/components"
	"sync"
)

var (
	globalClientMap = sync.Map{}
)

func Init(cm components.ConfigMap[*Config]) (func() error, error) {
	for k, v := range cm {
		Connect(k, v)
	}

	return IsHealth, nil
}

func Connect(configKey string, c *Config) {
	p, err := rocketmq.NewProducer(
		producer.WithNameServer([]string{c.NameServer}),
	)
	if err != nil {
		panic(err)
	}

	if err = p.Start(); err != nil {
		panic(err)
	}

	globalClientMap.Store(configKey, p)
}

func GetClient(ctx context.Context, keys ...string) rocketmq.Producer {
	configKey := "default"
	if len(keys) > 0 {
		configKey = keys[0]
	}

	if v, ok := globalClientMap.Load(configKey); ok {
		return v.(rocketmq.Producer)
	}

	panic(fmt.Sprintf("rocket mq producer %s not int", configKey))
}

func GetProducer[T any](ctx context.Context, topic string, keys ...string) *Producer[T] {
	return newProducer[T](GetClient(ctx, keys...), topic)
}

func IsHealth() (err error) {
	return err
}
