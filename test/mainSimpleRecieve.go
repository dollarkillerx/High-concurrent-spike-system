/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-16
* Time: 下午3:13
* */
package main

import (
	"High-concurrent-spike-system/RabbitMQ"
	"time"
)

func main() {
	for i:=0;i<10;i++ {
		go func() {
			simple := RabbitMQ.NewRabbitMQSimple("simpleTest")
			simple.ConsumeSimple()
		}()
	}
	for {
		time.Sleep(time.Second)
	}
}
