package queue

import "github.com/benedictotavio/payment_ms/pkg/message/rabbitmq"

type QueueConfig struct {
	ExchangeName string
	QueueName    string
	RoutingKey     string
}

func ConsumeQueue(
	queueConfig QueueConfig,
) {
	conn := rabbitmq.NewConnection()
	conn.CreateExchange(queueConfig.ExchangeName)
	conn.CreateQueue(
		queueConfig.QueueName,
		queueConfig.RoutingKey,
		queueConfig.ExchangeName,
	)
	conn.ConsumeQueue(queueConfig.QueueName)
}
