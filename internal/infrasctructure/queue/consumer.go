package queue

import "github.com/benedictotavio/payment_ms/pkg/message/rabbitmq"

func ConsumeQueue(queue string) {
	conn := rabbitmq.NewConnection()
	defer conn.Close()
	conn.ConsumeQueue(
		queue,
	)
}
