package main

import (
	"fmt"
	"promote/RabbitMQDemo/RabbitMQ"
	"strconv"
	"time"
)

// Work工作模式
func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" +
		"promoteSimple")

	for i := 0; i <= 100; i++ {
		rabbitmq.PublishSimple("Hello imooc!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
