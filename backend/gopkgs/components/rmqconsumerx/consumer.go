package rmqconsumerx

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
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
	p, err := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{c.NameServer}),
		consumer.WithGroupName(c.ConsumerGroup),
	)
	if err != nil {
		panic(err)
	}

	if err = p.Start(); err != nil {
		panic(err)
	}

	globalClientMap.Store(configKey, p)
}

func GetClient(ctx context.Context, keys ...string) rocketmq.PushConsumer {
	configKey := "default"
	if len(keys) > 0 {
		configKey = keys[0]
	}

	if v, ok := globalClientMap.Load(configKey); ok {
		return v.(rocketmq.PushConsumer)
	}

	panic(fmt.Sprintf("rocket mq consumer %s not int", configKey))
}

func GetConsumer[T any](ctx context.Context, topic string, keys ...string) *Consumer[T] {
	return newConsumer[T](GetClient(ctx, keys...), topic)
}

func IsHealth() (err error) {
	return err
}
