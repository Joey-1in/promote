package main

import (
	"fmt"
	"promote/RabbitMQDemo/RabbitMQ"
	"strconv"
	"time"
)

func main() {
	imoocOne := RabbitMQ.NewRabbitMQRouting("promoteo", "promoteOne")
	imoocTwo := RabbitMQ.NewRabbitMQRouting("promotet", "promoteTwo")
	for i := 0; i <= 10; i++ {
		imoocOne.PublishRouting("Hello promote one!" + strconv.Itoa(i))
		imoocTwo.PublishRouting("Hello promote Two!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}

}
