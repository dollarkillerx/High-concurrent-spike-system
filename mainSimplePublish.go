/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-16
* Time: 下午3:09
* */
package main

import (
	"High-concurrent-spike-system/RabbitMQ"
	"log"
)

func main() {
	simple := RabbitMQ.NewRabbitMQSimple("simpleTest")
	err := simple.PublishSimple("hello")
	if err != nil {
		log.Println(err.Error())
		log.Println("发送失败")
		return
	}

}
