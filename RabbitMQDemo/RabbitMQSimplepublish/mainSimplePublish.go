package main

import (
	"fmt"
	"promote/RabbitMQDemo/RabbitMQ"
)

// Simple模式
func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("promoteSimple")
	rabbitmq.PublishSimple("Hello Promote")
	fmt.Println("rabbitmq over")
}
