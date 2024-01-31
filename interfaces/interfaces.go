package interfaces

import "github.com/arturspolizel/payments/model"

type PaymentController interface {
	Get(uint) (model.Payment, error)
	Create(model.Payment) (uint, error)
}

type PaymentRepository interface {
	Get(uint) (model.Payment, error)
	Create(model.Payment) (uint, error)
}

type MerchantController interface {
	Get(uint) (model.Merchant, error)
	Create(model.Merchant) (uint, error)
}

type MerchantRepository interface {
	Get(uint) (model.Merchant, error)
	Create(model.Merchant) (uint, error)
}
