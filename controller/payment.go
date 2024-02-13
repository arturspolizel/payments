package controller

import (
	"slices"
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
	payment.Status = model.Captured
	id, err := c.paymentRepository.Create(payment)
	return id, err
}

func (c *PaymentController) Authorize(payment model.Payment) (uint, error) {
	payment.Total = payment.Amount + payment.Tips
	payment.Status = model.Authorized
	id, err := c.paymentRepository.Create(payment)
	return id, err
}

func (c *PaymentController) Capture(id uint, amount, tips int) error {
	payment, err := c.paymentRepository.Get(id)
	if err != nil {
		return err
	}

	if payment.Status != model.Authorized {
		err = &model.ErrInvalidPaymentStatus{
			Id:              id,
			AllowedStatuses: []model.PaymentStatus{model.Authorized},
		}
		return err
	}

	payment.Tips = tips
	payment.Amount = amount
	payment.Total = amount + tips

	err = c.paymentRepository.Update(payment)
	return err
}

func (c *PaymentController) Refund(id uint, amount, tips int) error {
	payment, err := c.paymentRepository.Get(id)
	if err != nil {
		return err
	}

	if slices.Contains(model.RefundAllowedStatuses, payment.Status) {
		err = &model.ErrInvalidPaymentStatus{
			Id:              id,
			AllowedStatuses: model.RefundAllowedStatuses,
		}
		return err
	}

	err = c.paymentRepository.Update(payment)
	return err
}

func (c *PaymentController) Void(id uint) error {
	payment, err := c.paymentRepository.Get(id)
	if err != nil {
		return err
	}

	if payment.Status != model.Authorized {
		err = &model.ErrInvalidPaymentStatus{
			Id:              id,
			AllowedStatuses: []model.PaymentStatus{model.Authorized},
		}
		return err
	}

	payment.Status = model.Void

	err = c.paymentRepository.Update(payment)
	return err
}
