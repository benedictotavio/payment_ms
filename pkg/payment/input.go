package payment

type ProductData struct {
	Name        string
	Description string
}

type PaymentInput struct {
	UserId      uint
	Amount      float64
	OrderId     uint
	ProductData ProductData
}
