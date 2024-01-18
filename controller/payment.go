package controller

import "github.com/arturspolizel/payments/interfaces"

type PaymentController struct {
	paymentRepository interfaces.PaymentRepository
}

func New(paymentRepository interfaces.PaymentRepository) *PaymentController {
	return &PaymentController{
		paymentRepository: paymentRepository,
	}
}
