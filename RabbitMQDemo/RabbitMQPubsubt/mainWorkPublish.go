package main

import (
	"promote/RabbitMQDemo/RabbitMQ"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQPubSub("" + "promoteSimple")
	rabbitmq.RecieveSub()
}
