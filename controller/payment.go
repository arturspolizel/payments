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

func (c *PaymentController) Get(id uint) model.Payment {
	payment := c.paymentRepository.Get(id)
	return payment
}

func (c *PaymentController) Create(payment model.Payment) uint {
	id := c.paymentRepository.Create(payment)
	return id
}
