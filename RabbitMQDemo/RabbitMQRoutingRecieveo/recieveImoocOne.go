package main

import (
	"promote/RabbitMQDemo/RabbitMQ"
)

func main() {
	imoocOne := RabbitMQ.NewRabbitMQRouting("promoteo", "promoteOne")
	imoocOne.RecieveRouting()
}
