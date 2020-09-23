package main

import (
	"fmt"
	"promote/RabbitMQDemo/RabbitMQ"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("promoteSimple")
	rabbitmq.ConsumeSimple()
	fmt.Println("rabbitmq over")
}
