package model

import (
	"fmt"
)

type ErrInvalidPaymentStatus struct {
	Id              uint
	AllowedStatuses []PaymentStatus
}

func (e *ErrInvalidPaymentStatus) Error() string {
	return fmt.Sprintf("Invalid status for operation on payment %v, allowed statuses %v", e.Id, e.AllowedStatuses)
}

type ErrInvalidTransactionAmount struct {
	Id     uint
	Amount int // Available amount for operation
	Tips   int // Available tips for operation
}

func (e *ErrInvalidTransactionAmount) Error() string {
	return fmt.Sprintf("Invalid amounts for operation on payment %v, available amount %v and tips %v", e.Id, e.Amount, e.Tips)
}
