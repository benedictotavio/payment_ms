package main

import (
	"os"

	"github.com/benedictotavio/payment_ms/internal/http"
	"github.com/benedictotavio/payment_ms/internal/infrasctructure/queue"
	"github.com/gin-gonic/gin"
)

func main() {
	go queue.ConsumeQueue(
		queue.QueueConfig{
			ExchangeName: "payment.service",
			QueueName:    "payments",
			RoutingKey:   "payment.receive",
		},
	)
	startServer()
}

func startServer() {
	gin := gin.Default()
	handler := http.PaymentHandler{}
	handler.RegisterRoutes(gin)

	var port = os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	gin.Run(":" + port)
}
