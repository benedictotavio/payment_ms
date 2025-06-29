package main

import (
	"github.com/benedictotavio/payment_ms/internal/http"
	"github.com/benedictotavio/payment_ms/internal/infrasctructure/queue"
	"github.com/gin-gonic/gin"
)

func main() {
	queue.ConsumeQueue("payments")
	gin := gin.Default()
	handler := http.PaymentHandler{}
	handler.RegisterRoutes(gin)
	gin.Run(":8080")
}
