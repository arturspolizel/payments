package model

import (
	"fmt"
)

type ErrInvalidUserStatus struct {
	Id              uint
	AllowedStatuses []UserStatus
}

func (e *ErrInvalidUserStatus) Error() string {
	return fmt.Sprintf("Invalid status for operation on user %v, allowed statuses %v", e.Id, e.AllowedStatuses)
}

type ErrInvalidEmailCode struct {
}

func (e *ErrInvalidEmailCode) Error() string {
	return fmt.Sprintf("Invalid validation code")
}

type ErrAuthenticationFailed struct {
}

func (e *ErrAuthenticationFailed) Error() string {
	return fmt.Sprintf("Couldn't authenticate user with provided credentials. Review username and password")
}
