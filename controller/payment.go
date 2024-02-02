package controller

import (
	"time"

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

func (c *PaymentController) List(startId, pageSize uint, startDate, endDate time.Time) ([]model.Payment, error) {
	payments, err := c.paymentRepository.List(startId, pageSize, startDate, endDate)
	return payments, err
}

func (c *PaymentController) Create(payment model.Payment) (uint, error) {
	payment.Total = payment.Amount + payment.Tips
	id, err := c.paymentRepository.Create(payment)
	return id, err
}
