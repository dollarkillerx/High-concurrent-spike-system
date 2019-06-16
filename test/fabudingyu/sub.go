/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-16
* Time: 下午5:39
* */
package main

import "High-concurrent-spike-system/RabbitMQ"

func main() {
	sub := RabbitMQ.NewRabbitMQPubSub("newProduct")
	sub.ConsumptionSub()
}