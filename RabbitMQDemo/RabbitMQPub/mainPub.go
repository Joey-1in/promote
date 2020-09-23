package main

import (
	"fmt"
	"promote/RabbitMQDemo/RabbitMQ"
	"strconv"
	"time"
)

// 订阅模式
func main() {
	rabbitmq := RabbitMQ.NewRabbitMQPubSub("" +
		"promoteSimple")
	for i := 0; i < 100; i++ {
		rabbitmq.PublishPub("订阅模式生产第" +
			strconv.Itoa(i) + "条" + "数据")
		fmt.Println("订阅模式生产第" +
			strconv.Itoa(i) + "条" + "数据")
		time.Sleep(1 * time.Second)
	}
}
