/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-16
* Time: 下午3:13
* */
package main

import "High-concurrent-spike-system/RabbitMQ"

func main() {
	simple := RabbitMQ.NewRabbitMQSimple("simpleTest")
	simple.ConsumeSimple()
}
