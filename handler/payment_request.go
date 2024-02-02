package handler

import (
	"time"

	"github.com/arturspolizel/payments/model"
)

type PaymentCreateRequest struct {
	MerchantId uint           `json:"merchantId" binding:"required"`
	Amount     int            `json:"amount" binding:"required"`
	Tips       int            `json:"tips" binding:"required"`
	Currency   model.Currency `json:"currency" binding:"required"`
}

type PaymentListRequest struct {
	PaginationRequest
	StartDate time.Time `json:"startDate" form:"startDate" binding:"required"`
	EndDate   time.Time `json:"endDate" form:"endDate" binding:"required"`
}

func (pcr *PaymentCreateRequest) toPayment() model.Payment {
	payment := model.Payment{}
	payment.Amount = pcr.Amount
	payment.Tips = pcr.Tips
	payment.MerchantId = pcr.MerchantId
	payment.Currency = pcr.Currency

	return payment
}
