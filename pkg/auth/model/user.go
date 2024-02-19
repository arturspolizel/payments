package model

import "time"

type User struct {
	ID           uint       `json:"id"`
	MerchantId   uint       `json:"merchantId"`
	Name         string     `json:"name"`
	Email        string     `json:"email"`
	PasswordHash string     `json:"passwordHash"`
	Status       UserStatus `json:"status"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
}

type UserStatus string

const (
	PendingActivation UserStatus = "PendingActivation"
	Active            UserStatus = "Active"
	Locked            UserStatus = "Locked"
)

var userStatuses = map[UserStatus]bool{
	PendingActivation: true,
	Active:            true,
	Locked:            true,
}

func (c UserStatus) Validate() bool {
	return userStatuses[c]
}

type ValidationEmail struct {
	ID        uint      `json:"id"`
	UserId    uint      `json:"userId"`
	Code      string    `json:"code"`
	Validated bool      `json:"validated"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
