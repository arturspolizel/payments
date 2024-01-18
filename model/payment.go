package model

import "time"

type Payment struct {
	ID         uint
	MerchantId string
	Amount     int
	Tips       int
	Total      int
	Currency   Currency
	CreatedAt  time.Time // Automatically managed by GORM for creation time
	UpdatedAt  time.Time
}

// Weekday - Custom type to hold value for weekday ranging from 1-7
type Currency string

const (
	USD Currency = "USD"
	BRL Currency = "BRL"
	EUR Currency = "EUR"
)
