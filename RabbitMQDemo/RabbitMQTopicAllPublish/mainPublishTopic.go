package main

import (
	"fmt"
	"promote/RabbitMQDemo/RabbitMQ"
	"strconv"
	"time"
)

func main() {
	imoocOne := RabbitMQ.NewRabbitMQTopic("promoteTopic", "promote.topic.one")
	imoocTwo := RabbitMQ.NewRabbitMQTopic("promoteTopic", "promote.topic.two")
	for i := 0; i <= 10; i++ {
		imoocOne.PublishTopic("Hello promote topic one!" + strconv.Itoa(i))
		imoocTwo.PublishTopic("Hello promote topic Two!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}

}
