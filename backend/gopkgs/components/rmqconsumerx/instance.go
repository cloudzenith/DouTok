package rmqconsumerx

import (
	"context"
	"encoding/json"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

type Consumer[T any] struct {
	consumer rocketmq.PushConsumer
	topic    string
}

func newConsumer[T any](consumer rocketmq.PushConsumer, topic string) *Consumer[T] {
	return &Consumer[T]{
		consumer: consumer,
		topic:    topic,
	}
}

func (c *Consumer[T]) GetInstance() rocketmq.PushConsumer {
	return c.consumer
}

func (c *Consumer[T]) Subscribe(f func(context.Context, T, *primitive.MessageExt) (consumer.ConsumeResult, error)) error {
	return c.consumer.Subscribe(c.topic, consumer.MessageSelector{}, func(ctx context.Context, messages ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		result := consumer.ConsumeSuccess
		for _, msg := range messages {
			data := new(T)
			err := json.Unmarshal(msg.Body, data)
			if err != nil {
				return consumer.Rollback, err
			}

			r, err := f(ctx, *data, msg)
			if err != nil {
				return consumer.Rollback, err
			}

			result &= r
		}

		return result, nil
	})
}

func (c *Consumer[T]) SubscribeWithSelector(selector consumer.MessageSelector, f func(context.Context, T, ...*primitive.MessageExt) (consumer.ConsumeResult, error)) error {
	return c.consumer.Subscribe(c.topic, selector, func(ctx context.Context, messages ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		result := consumer.ConsumeSuccess
		for _, msg := range messages {
			data := new(T)
			err := json.Unmarshal(msg.Body, data)
			if err != nil {
				return consumer.Rollback, err
			}

			r, err := f(ctx, *data, msg)
			if err != nil {
				return consumer.Rollback, err
			}

			result &= r
		}

		return result, nil
	})
}

func (c *Consumer[T]) Unsubscribe() error {
	return c.consumer.Unsubscribe(c.topic)
}

func (c *Consumer[T]) Suspend() {
	c.consumer.Suspend()
}

func (c *Consumer[T]) Resume() {
	c.consumer.Resume()
}

func (c *Consumer[T]) GetOffsetDiffMap() map[string]int64 {
	return c.consumer.GetOffsetDiffMap()
}
