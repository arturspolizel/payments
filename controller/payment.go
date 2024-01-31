package controller

import (
	"github.com/arturspolizel/payments/interfaces"
	"github.com/arturspolizel/payments/model"
)

type PaymentController struct {
	paymentRepository interfaces.PaymentRepository
}

func NewPaymentController(paymentRepository interfaces.PaymentRepository) *PaymentController {
	return &PaymentController{
		paymentRepository: paymentRepository,
	}
}

func (c *PaymentController) Get(id uint) (model.Payment, error) {
	payment, err := c.paymentRepository.Get(id)
	return payment, err
}

func (c *PaymentController) Create(payment model.Payment) (uint, error) {
	payment.Total = payment.Amount + payment.Tips
	id, err := c.paymentRepository.Create(payment)
	return id, err
}
