package rabbitmq

import (
	"fmt"
	"log"
	"os"
	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQConnection struct {
	Conn    *amqp091.Connection
	Channel *amqp091.Channel
}

func NewConnection() *RabbitMQConnection {
	conn, err := amqp091.Dial(
		os.Getenv("RABBITMQ_URL"),
	)

	if err != nil {
		panic(err)
	}

	channel, err := conn.Channel()

	if err != nil {
		panic(err)
	}

	return &RabbitMQConnection{
		Conn:    conn,
		Channel: channel,
	}
}

func (r *RabbitMQConnection) Close() error {
	return r.Channel.Close()
}

func (r *RabbitMQConnection) ConsumeQueue(queue string) (string, error) {

	var message string

	defer r.Conn.Close()

	msgs, err := r.Channel.Consume(
		queue, // queue
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)

	if err != nil {
		panic(err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			fmt.Printf("Received a message: %s\n", string(d.Body))
			message = string(d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	<-forever
	return message, nil
}

func (r *RabbitMQConnection) CreateExchange(exchangeName string) (string, error) {
	err := r.Channel.ExchangeDeclare(
		exchangeName,
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)

	if err != nil {
		return "", err
	}

	return "", nil
}

func (r *RabbitMQConnection) CreateQueue(queueName string, routingKey string, exchange string) error {
	_, err := r.Channel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)

	if err != nil {
		return err
	}

	err = r.Channel.QueueBind(
		queueName,  // queue
		routingKey, // routing key
		exchange,   // exchange
		false,      // no-wait
		nil,        // arguments
	)

	if err != nil {
		return err
	}

	return nil
}