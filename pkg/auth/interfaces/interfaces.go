package interfaces

import (
	"github.com/arturspolizel/payments/pkg/auth/model"
)

type UserController interface {
	Get(uint) (model.User, error)
	Create(model.User, string) (uint, error)
	Login(user, password string) (string, error)
	Validate(code string) error
}

type UserRepository interface {
	Get(uint) (model.User, error)
	GetByEmail(email string) (model.User, error)
	Update(model.User) error
	Create(model.User) (uint, error)
	CreateValidationEmail(model.ValidationEmail) (uint, error)
	GetEmailByCode(code string) (model.ValidationEmail, error)
	UpdateValidationEmail(model.ValidationEmail) error
}

type EmailAdapter interface {
	SendEmail(address, content string) (err error)
}
