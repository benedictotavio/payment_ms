package main

import (
	"github.com/benedictotavio/payment_ms/internal/http"
	"github.com/benedictotavio/payment_ms/pkg/message/rabbitmq"
	"github.com/gin-gonic/gin"
)

func main() {
	rabbitmq := rabbitmq.NewConnection()
	defer rabbitmq.Close()
	go rabbitmq.ConsumeQueue("payments")

	gin := gin.Default()
	handler := http.PaymentHandler{}
	handler.RegisterRoutes(gin)
	gin.Run(":8080")
}
