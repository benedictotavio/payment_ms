package queue

import "github.com/benedictotavio/payment_ms/pkg/message/rabbitmq"


type Publisher struct {
	producer *rabbitmq.RabbitMQProducer
}

func (p *Publisher) Publish(queue string, body []byte) error {
	producer := rabbitmq.RabbitMQProducer{
		Channel: p.producer.Channel,
	}
	return producer.Publish(queue, body, "")
}