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
	"strconv"
)

func main() {
	simple := RabbitMQ.NewRabbitMQSimple("simpleTest")
	for i:=0;i<999999;i++ {
		err := simple.PublishSimple("hello: " + strconv.Itoa(i))
		if err != nil {
			log.Println(err.Error())
			log.Println("发送失败")
			return
		}
	}
}
