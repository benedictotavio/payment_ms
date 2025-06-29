package rabbitmq

import (
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQProoducer struct {
	Channel *amqp091.Channel
}

func (r *RabbitMQProoducer) Publish(queue string, body []byte, exchange string) error {

	q, err := r.Channel.QueueDeclare(
		queue,
		false,
		false,
		false,
		false,
		amqp091.Table{
			exchange: exchange,
			"x-dead-letter-exchange": "payment.service",
			"x-dead-letter-routing-key": "payments",
			"x-message-ttl": 10000,
			"x-max-length": 10000,
		},
	)

	if err != nil {
		return err
	}

	err = r.Channel.Publish(
		exchange, // exchange
		q.Name, // routing key
		false, // mandatory
		false, // immediate
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        body,
			Timestamp: time.Now(),
			Expiration: "10000",
		},
	)

	if err != nil {
		return err
	}

	log.Printf(" [x] Sent %s", body)

	return nil
}