package main

import (
	"promote/RabbitMQDemo/RabbitMQ"
)

func main() {
	imoocOne := RabbitMQ.NewRabbitMQTopic("promoteTopic", "#")
	imoocOne.RecieveTopic()
}
