package main

import (
	"promote/RabbitMQDemo/RabbitMQ"
)

func main() {
	imoocOne := RabbitMQ.NewRabbitMQRouting("promotet", "promoteTwo")
	imoocOne.RecieveRouting()
}
