/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-16
* Time: 下午6:04
* */
package main

import "High-concurrent-spike-system/RabbitMQ"

func main() {
	two := RabbitMQ.NewRabbitMQPubRoutiong("exOne","ones")

	//two := RabbitMQ.NewRabbitMQPubRoutiong("exOne","twos")
	two.ConsumptionRouting()
}