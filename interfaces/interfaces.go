package interfaces

import "github.com/arturspolizel/payments/model"

type PaymentController interface {
	Get(id string) model.Payment
	Create(model.Payment) string
}

type PaymentRepository interface {
}
