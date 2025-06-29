package createpayment

import (
	"context"
	"github.com/benedictotavio/payment_ms/internal/domain"
	"github.com/benedictotavio/payment_ms/internal/infrasctructure/db"
)

type CreatePaymentUsecase interface {
	CreateUser(input CreatePaymentInput) (CreatePaymentOutput, error)
}

type createPaymentUsecase struct {
	repository db.PaymentRepository
	ctx        context.Context
}

func NewCreatePaymentUsecase(ctx context.Context, repository db.PaymentRepository) *createPaymentUsecase {
	return &createPaymentUsecase{repository: repository, ctx: ctx}
}

func (uc *createPaymentUsecase) CreateUser(user CreatePaymentInput) error {
	if err := uc.repository.Create(
		domain.Payment{
			ID:     user.OrderId,
			Amount: 0,
			Method: "",
		},
	); err != nil {
		return err
	}
	return nil
}

func (uc *createPaymentUsecase) GetUser(id int) (domain.Payment, error) {
	user, err := uc.repository.Get(id)
	if err != nil {
		return domain.Payment{}, err
	}
	return user, nil
}
