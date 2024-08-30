package rmqproducerx

import (
	"context"
	"encoding/json"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type Producer[T any] struct {
	producer rocketmq.Producer
	topic    string
}

func newProducer[T any](producer rocketmq.Producer, topic string) *Producer[T] {
	return &Producer[T]{
		producer: producer,
		topic:    topic,
	}
}

func (p *Producer[T]) GetInstance() rocketmq.Producer {
	return p.producer
}

type ProduceOption func(*primitive.Message) *primitive.Message

func WithProperties(properties map[string]string) ProduceOption {
	return func(message *primitive.Message) *primitive.Message {
		message.WithProperties(properties)
		return message
	}
}

func WithProperty(key string, value string) ProduceOption {
	return func(message *primitive.Message) *primitive.Message {
		message.WithProperty(key, value)
		return message
	}
}

func WithDelayTimeLevel(level int) ProduceOption {
	return func(message *primitive.Message) *primitive.Message {
		return message.WithDelayTimeLevel(level)
	}
}

func WithTag(tag string) ProduceOption {
	return func(message *primitive.Message) *primitive.Message {
		return message.WithTag(tag)
	}
}

func WithKeys(keys ...string) ProduceOption {
	return func(message *primitive.Message) *primitive.Message {
		return message.WithKeys(keys)
	}
}

func WithShardingKey(key string) ProduceOption {
	return func(message *primitive.Message) *primitive.Message {
		return message.WithShardingKey(key)
	}
}

func (p *Producer[T]) marshalMessage(message T, options ...ProduceOption) (*primitive.Message, error) {
	data, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}

	msg := primitive.NewMessage(p.topic, data)
	for _, option := range options {
		msg = option(msg)
	}

	return msg, nil
}

func (p *Producer[T]) SendSync(ctx context.Context, message T, options ...ProduceOption) (*primitive.SendResult, error) {
	msg, err := p.marshalMessage(message, options...)
	if err != nil {
		return nil, err
	}

	return p.producer.SendSync(ctx, msg)
}

func (p *Producer[T]) SendAsync(ctx context.Context, message T, options ...ProduceOption) error {
	msg, err := p.marshalMessage(message, options...)
	if err != nil {
		return err
	}

	return p.producer.SendAsync(ctx, defaultSendAsyncCallback, msg)
}

func (p *Producer[T]) SendAsyncWithCallback(ctx context.Context, message T, callback func(ctx context.Context, result *primitive.SendResult, err error), options ...ProduceOption) error {
	msg, err := p.marshalMessage(message, options...)
	if err != nil {
		return err
	}

	return p.producer.SendAsync(ctx, callback, msg)
}

func (p *Producer[T]) SendOneWay(ctx context.Context, message T, options ...ProduceOption) error {
	msg, err := p.marshalMessage(message, options...)
	if err != nil {
		return err
	}

	return p.producer.SendOneWay(ctx, msg)
}

func (p *Producer[T]) Request(ctx context.Context, message T, ttl time.Duration, options ...ProduceOption) (*primitive.Message, error) {
	msg, err := p.marshalMessage(message, options...)
	if err != nil {
		return nil, err
	}

	return p.producer.Request(ctx, ttl, msg)
}

func (p *Producer[T]) RequestAsync(ctx context.Context, callback func(ctx context.Context, msg *primitive.Message, err error), message T, ttl time.Duration, options ...ProduceOption) error {
	msg, err := p.marshalMessage(message, options...)
	if err != nil {
		return err
	}

	return p.producer.RequestAsync(ctx, ttl, callback, msg)
}

func defaultSendAsyncCallback(ctx context.Context, result *primitive.SendResult, err error) {
	var level log.Level
	var msg string
	if err != nil {
		level = log.LevelError
		msg = "failed to send message"
	} else {
		level = log.LevelInfo
		msg = "message sent"
	}

	log.Context(ctx).Log(
		level,
		"msg", msg,
		"status", result.Status,
		"msg_id", result.MsgID,
		"topic", result.MessageQueue.Topic,
		"broker", result.MessageQueue.BrokerName,
		"queue_id", result.MessageQueue.QueueId,
		"queue_offset", result.QueueOffset,
		"transaction_id", result.TransactionID,
		"offset_msg_id", result.OffsetMsgID,
		"region_id", result.RegionID,
		"trace_on", result.TraceOn,
	)
}
