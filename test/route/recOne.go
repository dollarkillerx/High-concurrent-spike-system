/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-16
* Time: 下午6:03
* */
package main

import "High-concurrent-spike-system/RabbitMQ"

func main() {
	one := RabbitMQ.NewRabbitMQPubRoutiong("exOne","ones")
	one.ConsumptionRouting()
}
