package v1

import (
	"box/library"
	"box/rabbitmq"
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

type Order struct {
	//ID        int     `json:"id"`
	Name      string  `json:"name"`
	Price     float32 `json:"price"`
	Number    int     `json:"number"`
	UserId    int     `json:"user_id"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

func SaveOrder(c *gin.Context) {

	responseBody := library.NewResponseBody()
	defer library.RecoverResponse(c, responseBody)
	order := Order{}
	queueExchange := &rabbitmq.QueueExchange{
		"simple",
		"simple-queue",
		"simple-queue",
		"",
		"",
	}
	mq := rabbitmq.New(queueExchange)
	c.Bind(order)
	data, err := json.Marshal(order)
	if err != nil {
		// TODO:: 应该是非阻塞
		responseBody.SetMessage("数据错误")
	}
	mq.RegisterProducer(data)
	// 启动
	mq.Start()

	responseBody.SetCode(0)
	responseBody.SetMessage("发送成功")

	return

}

const MQURL = "amqp://admin:admin123@127.0.0.1:5672/box"

type RabbitMQ struct {
	//连接
	conn *amqp.Connection
	//管道
	channel *amqp.Channel
	//队列名称
	QueueName string
	//交换机
	Exchange string
	//key Simple模式 几乎用不到
	Key string
	//连接信息
	Mqurl string
}

func NewRabbitMQ(queuename string, exchange string, key string) *RabbitMQ {
	rabbitmq := &RabbitMQ{QueueName: queuename, Exchange: exchange, Key: key, Mqurl: MQURL}
	var err error
	//创建rabbitmq连接
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	if err != nil {
		log.Println("连接失败", err)
	}
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	if err != nil {
		log.Println("获取channel失败", err)
	}
	rabbitmq.channel.QueueDeclare(queuename, true, false, false, true, nil)
	return rabbitmq
}

type ExchangeConfig struct {
	FanoutExchange string
	SendMsmQueue   string
	SendMailQueue  string
	Kind           string
}

func SaveOrderAction(c *gin.Context) {

	ec := &ExchangeConfig{
		FanoutExchange: "fanout_exchange",
		SendMsmQueue:   "fanout_msm_queue",
		SendMailQueue:  "fanout_mail_queue",
		Kind:           "fanout",
	}

	rabbitmq_msm := &RabbitMQ{QueueName: ec.SendMsmQueue, Exchange: ec.FanoutExchange, Key: ec.SendMsmQueue, Mqurl: MQURL}
	var err error
	//创建rabbitmq连接
	rabbitmq_msm.conn, err = amqp.Dial(rabbitmq_msm.Mqurl)
	if err != nil {
		log.Println("连接失败", err)
	}

	rabbitmq_msm.channel, err = rabbitmq_msm.conn.Channel()
	if err != nil {
		log.Println("获取channel失败", err)
	}
	err = rabbitmq_msm.channel.Confirm(false)
	if err != nil {
		log.Println("发布确认~~~", err)
	}

	// 交换机
	err = rabbitmq_msm.channel.ExchangeDeclare(ec.FanoutExchange, ec.Kind, true, false, false, false, nil)

	if err != nil {
		log.Println("交换机声明失败", err)
	}

	rabbitmq_msm.channel.QueueDeclare(ec.SendMsmQueue, true, false, false, true, nil)
	rabbitmq_msm.channel.QueueDeclare(ec.SendMailQueue, true, false, false, true, nil)
	// 绑定

	rabbitmq_msm.channel.QueueBind(ec.SendMsmQueue, ec.SendMsmQueue, ec.FanoutExchange, false, nil)
	rabbitmq_msm.channel.QueueBind(ec.SendMailQueue, ec.SendMsmQueue, ec.FanoutExchange, false, nil)

	msg := "发短信"
	rabbitmq_msm.channel.Publish(ec.FanoutExchange, ec.SendMsmQueue, true, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(msg),
	})

}
