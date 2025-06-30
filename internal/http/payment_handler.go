package http

import (
	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
}

func (p *PaymentHandler) RegisterRoutes(r *gin.Engine) {
	r.GET("/payments", p.ListPayments)
}

func (p *PaymentHandler) ListPayments(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
