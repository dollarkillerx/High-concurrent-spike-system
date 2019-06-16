/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-16
* Time: 下午6:01
* */
package main

import (
	"High-concurrent-spike-system/RabbitMQ"
	"strconv"
	"time"
)

func main() {
	one := RabbitMQ.NewRabbitMQPubRoutiong("exOne","ones")
	two := RabbitMQ.NewRabbitMQPubRoutiong("exOne","twos")

	for i:=0 ;i<30;i++ {
		one.PublishRouting("Hello one!" + strconv.Itoa(i))
		two.PublishRouting("Hello two!" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}