/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-16
* Time: 下午5:41
* */
package main

import (
	"High-concurrent-spike-system/RabbitMQ"
	"strconv"
)

// 生产者

func main() {
	sub := RabbitMQ.NewRabbitMQPubSub("newProduct")
	for i:=0;i<1000;i++ {
		sub.PublishPub("订阅模式生产第: " + strconv.Itoa(i))
	}
}