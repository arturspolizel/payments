package model

import "time"

type Payment struct {
	ID         uint      `json:"id"`
	MerchantId string    `json:"merchantId"`
	Amount     int       `json:"amount"`
	Tips       int       `json:"tips"`
	Total      int       `json:"total"`
	Currency   Currency  `json:"currency"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
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
