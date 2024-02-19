package model

import "time"

type Payment struct {
	ID         uint          `json:"id"`
	MerchantId uint          `json:"merchantId"`
	Merchant   Merchant      `json:"merchant"`
	Amount     int           `json:"amount"`
	Tips       int           `json:"tips"`
	Total      int           `json:"total"`
	Status     PaymentStatus `json:"paymentStatus"`
	Method     PaymentMethod `json:"paymentMethod"`
	Refunds    []Refund      `json:"refunds"`
	Currency   Currency      `json:"currency"`
	CreatedAt  time.Time     `json:"createdAt"`
	UpdatedAt  time.Time     `json:"updatedAt"`
}

type Refund struct {
	ID        uint `json:"id"`
	PaymentId uint `json:"paymentId"`
	Amount    int  `json:"amount"`
	Tips      int  `json:"tips"`
	Total     int  `json:"total"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

var RefundAllowedStatuses = []PaymentStatus{
	Captured,
}

func (p *Payment) GetRefundableAmount() (int, int) {
	refundableAmount := p.Amount
	refundableTips := p.Tips
	for _, refund := range p.Refunds {
		refundableAmount -= refund.Amount
		refundableTips -= refund.Tips
	}

	return refundableAmount, refundableTips
}

type Currency string

const (
	USD Currency = "USD"
	BRL Currency = "BRL"
	EUR Currency = "EUR"
)

var currencies = map[Currency]bool{
	USD: true,
	BRL: true,
	EUR: true,
}

func (c Currency) Validate() bool {
	return currencies[c]
}

type PaymentStatus string

const (
	Authorized PaymentStatus = "authorized"
	Captured   PaymentStatus = "captured"
	Void       PaymentStatus = "void"
)

var paymentStatuses = map[PaymentStatus]bool{
	Authorized: true,
	Captured:   true,
	Void:       true,
}

func (ps PaymentStatus) Validate() bool {
	return paymentStatuses[ps]
}

type PaymentMethod string

const (
	CardPresent    PaymentMethod = "cardpresent"
	CardNotPresent PaymentMethod = "cardnotpresent"
	Cash           PaymentMethod = "cash"
)

var paymentMethods = map[PaymentMethod]bool{
	CardPresent:    true,
	CardNotPresent: true,
	Cash:           true,
}

func (pm PaymentMethod) Validate() bool {
	return paymentMethods[pm]
}
