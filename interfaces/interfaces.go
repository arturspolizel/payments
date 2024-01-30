package interfaces

import "github.com/arturspolizel/payments/model"

type PaymentController interface {
	Get(id uint) model.Payment
	Create(payment model.Payment) uint
}

type PaymentRepository interface {
	Get(id uint) model.Payment
	Create(payment model.Payment) uint
}

type MerchantController interface {
	Get(id uint) model.Merchant
	Create(payment model.Merchant) uint
}

type MerchantRepository interface {
	Get(id uint) model.Merchant
	Create(payment model.Merchant) uint
}
