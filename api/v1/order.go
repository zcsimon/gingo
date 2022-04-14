package v1

import (
	"box/library"
	"box/rabbitmq"
	"encoding/json"

	"github.com/gin-gonic/gin"
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
	queueName := make([]string, 2, 4)
	queueName[0] = "simple-queue"
	queueExchange := &rabbitmq.QueueExchange{
		"simple",
		queueName,
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

func Topic(c *gin.Context) {
	responseBody := library.NewResponseBody()
	defer library.RecoverResponse(c, responseBody)
	order := Order{}
	queueName := make([]string, 2, 4)
	queueName[0] = "topic-cluster-queue-msm"
	queueName[1] = "topic-cluster-queue-mail"
	queueExchange := &rabbitmq.QueueExchange{
		"topic",
		queueName,
		"topic",
		"topic-exchange",
		"topic",
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
	//

	responseBody.SetCode(0)
	responseBody.SetMessage("发送成功")

	return
}
