package main

import (
	"promote/RabbitMQDemo/RabbitMQ"
)

func main() {
	imoocOne := RabbitMQ.NewRabbitMQTopic("promoteTopic", "promote.*.two")
	imoocOne.RecieveTopic()
}
