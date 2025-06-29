package payment

import (
	"errors"
	"log"

	"github.com/stripe/stripe-go/v82/checkout/session"
	"github.com/stripe/stripe-go/v82"
)

type PaymentClient interface {
	SendPayment(input *PaymentInput) (*stripe.CheckoutSession, error)
	GetPaymentStatus(paymentId string) (*stripe.CheckoutSession, error)
}

type payment struct {
	stripeClientKey string
	successUrl      string
	failureUrl      string
}

// GetPaymentStatus implements PaymentClient.
func (p *payment) GetPaymentStatus(paymentId string) (*stripe.CheckoutSession, error) {
	stripe.Key = p.stripeClientKey
	session, err := session.Get(paymentId, nil)
	if err != nil {
		log.Println(err)
		return nil, errors.New("payment not found")
	}
	return session, nil
}

func (p *payment) SendPayment(input *PaymentInput) (*stripe.CheckoutSession, error) {
	stripe.Key = p.stripeClientKey
	amountInCents := int64(input.Amount * 100)
	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String("brl"),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String("Payment"),
						// Products are created in the Stripe dashboard.
					},
					UnitAmount: stripe.Int64(amountInCents),
				},
				Quantity: stripe.Int64(1),
			},
		},
		Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
		CancelURL: stripe.String(p.failureUrl),
		SuccessURL: stripe.String(p.successUrl),
	}

	params.AddMetadata("order_id", string(input.OrderId))
	params.AddMetadata("user_id", string(input.UserId))

	session, err := session.New(params)

	if err != nil {
		log.Printf("session.New: %v", err)
		return nil, errors.New("error creating stripe session")
	}

	return session, nil
}

func NewPaymentClient(stripeClientKey, successUrl, failureUrl string) PaymentClient {
	return &payment{
		stripeClientKey: stripeClientKey,
		successUrl:      successUrl,
		failureUrl:      failureUrl,
	}
}
