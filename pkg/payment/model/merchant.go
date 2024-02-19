package model

import "time"

type Merchant struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Active bool   `json:"active"`

	MaximumTransactionValue *uint     `json:"maximumTransactionValue,omitempty"`
	CreatedAt               time.Time `json:"createdAt"`
	UpdatedAt               time.Time `json:"updatedAt"`
}

// type Currency string

// const (
// 	USD Currency = "USD"
// 	BRL Currency = "BRL"
// 	EUR Currency = "EUR"
// )

// var currencies = map[Currency]bool{
// 	USD: true,
// 	BRL: true,
// 	EUR: true,
// }

// func (c Currency) Validate() bool {
// 	return currencies[c]
// }
