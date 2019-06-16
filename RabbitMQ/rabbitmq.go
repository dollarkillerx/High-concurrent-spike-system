/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-16
* Time: 下午2:02
* */
package RabbitMQ

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

// url 格式 amqp://账户:秘密@rabbitmq服务器地址:端口号/host
const MQURL  = "amqp://guest:guest@127.0.0.1:5672/"

type RabbitMQ struct {
	conn *amqp.Connection
	channel *amqp.Channel
	// 队列名称
	QueueName string
	// 交换机
	Exchange string
	// key
	Key string
	// 连接信息
	Mqurl string
}

// 创建RabbitMQ结构体实例
func NewRabbitMQ(queueName,exchange,key string) *RabbitMQ {
	mq := &RabbitMQ{
		QueueName: queueName,
		Exchange:  exchange,
		Key:       key,
		Mqurl:     MQURL,
	}
	connection, e := amqp.Dial(mq.Mqurl)
	mq.failOnErr(e,"创建连接失败")
	channel, e := connection.Channel()
	mq.failOnErr(e,"获取channel失败")
	mq.conn = connection
	mq.channel = channel
	return mq
}

// 断开channel和connection
func (r *RabbitMQ) Destroy() {
	r.channel.Close()
	r.conn.Close()
}

// 错误处理函数
func (r *RabbitMQ) failOnErr (err error,message string) {
	if err != nil {
		log.Fatal("%s:%s",message,err)
		panic(fmt.Sprintf("%s:%s",message,err))
	}
}

// 简单模式Step1: 创建简单模式下RabbitMQ
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	// 使用default exchange
	mq := NewRabbitMQ(queueName,"","")
	return mq
}

// 简单模式Step2: 简单模式下生产代码
func (r *RabbitMQ) PublishSimple(message string) error {
	// 1. 申请队列,如果队列不存在会自动创建,如果存在则跳过创建
	// 保证队列存在,消息能够发送到队列中
	_, e := r.channel.QueueDeclare(
		r.QueueName, // 队列名称
		false,       // 消息是否持久化 false不
		false,       // 是否自动删除  没有消息就删除次队列
		false,       // 是否具有排他性 (true只有自己可以访问,连接端口会自动删除)
		false,       // 是否阻塞  (设置为true时就像没有bufio的channel)
		nil,         // 额外的属性
	)
	if e != nil {
		log.Println(e.Error())
		return e
	}
	// 2. 发送消息到队列中
	e = r.channel.Publish(
		r.Exchange,// 交换机 名称
		r.QueueName, // 队列名称
		false,// 如果为true,根据exchange类型和routkey规则,如果无法找到符合条件的队列,那末就会吧发送的消息返回给发送者
		false,// [最新版以及失效这个参数]如果为true,但exchange发送消息到队列后发现队列上没有绑定消费者,则会吧消息发送给发送者
		amqp.Publishing{
			ContentType:"text/plain",// 明文
			Body:[]byte(message),
		},
	)
	return e
}

// 简单模式Step3: 消费消息
func (r *RabbitMQ) ConsumeSimple() error {
	// 1. 申请队列,如果队列不存在会自动创建,如果存在则跳过创建
	// 保证队列存在,消息能够发送到队列中
	_, e := r.channel.QueueDeclare(
		r.QueueName, // 队列名称
		false,       // 消息是否持久化 false不
		false,       // 是否自动删除  没有消息就删除次队列
		false,       // 是否具有排他性 (true只有自己可以访问)
		false,       // 是否阻塞  (设置为true时就像没有bufio的channel)
		nil,         // 额外的属性
	)
	if e != nil {
		log.Println(e.Error())
		return e
	}

	// 接受消息
	msgch, e := r.channel.Consume(
		r.QueueName, // 队列名称
		"", // 当前消费者名称 (用于区分多个消费者)
		true, // 是否自动应答
		false,// 是否具有排他性 (true只有自己可以访问,连接端口会自动删除)
		false, // [已经不支持] 如果设置为true,表示不能了个将同一个connection中发送的消息传递个ie这个connection中的消费者
		false, // 消费是否阻塞
		nil,// 其他参数
	)
	if e != nil {
		return e
	}

	foreve := make(chan bool)
	// 启动协程处理消息
	go func() {
		for {
			select {
			case data := <-msgch :
				// 实现消息处理
				bytes := data.Body
				fmt.Printf("%s",bytes)
			}
		}
	}()

	<-foreve
	return nil
}
