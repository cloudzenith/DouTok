package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"time"
)

func main() {
	c, _ := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{"127.0.0.1:9876"}),
		consumer.WithGroupName("test-group"),
	)
	err := c.Subscribe("test", consumer.MessageSelector{},
		func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
			for i := range msgs {
				fmt.Printf("获取到的数据： %v  \n", msgs[i])
			}
			return consumer.ConsumeSuccess, nil
		})
	if err != nil {
		fmt.Println("读取消息失败")
	}
	_ = c.Start()
	//不能让主 goroutine 退出，不然就马上结束了
	time.Sleep(time.Hour)
	_ = c.Shutdown()
}
