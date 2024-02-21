package interfaces

import (
	"time"

	"github.com/arturspolizel/payments/pkg/payment/model"
	"github.com/arturspolizel/payments/utils"
)

type PaymentController interface {
	Get(uint) (model.Payment, error)
	List(startId, pageSize uint, startDate, endDate time.Time) ([]model.Payment, error)
	Create(model.Payment) (uint, error)
	Authorize(model.Payment) (uint, error)
	Capture(id uint, amount, tips int) error
	Refund(id uint, amount, tips int) error
	Void(id uint) error
}

type PaymentRepository interface {
	Get(uint) (model.Payment, error)
	Update(model.Payment) error
	List(startId, pageSize uint, startDate, endDate time.Time) ([]model.Payment, error)
	Create(model.Payment) (uint, error)
	CreateRefund(refund model.Refund) (uint, error)
}

type MerchantController interface {
	Get(uint) (model.Merchant, error)
	List(startId, pageSize uint) ([]model.Merchant, error)
	Create(model.Merchant) (uint, error)
}

type MerchantRepository interface {
	Get(uint) (model.Merchant, error)
	List(startId, pageSize uint) ([]model.Merchant, error)
	Create(model.Merchant) (uint, error)
}

type JwtProcessor interface {
	NewToken(utils.TokenContext) (string, error)

	Validate(string) (utils.TokenContext, error)
}
