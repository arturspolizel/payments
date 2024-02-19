package interfaces

import (
	"github.com/arturspolizel/payments/pkg/auth/model"
)

type UserController interface {
	Get(uint) (model.User, error)
	Create(model.User, string) (uint, error)
	Login(user, password string) (model.User, error)
	Lock(id uint) error
}

type UserRepository interface {
	Get(uint) (model.User, error)
	GetByEmail(email string) (model.User, error)
	Update(model.User) error
	Create(model.User) (uint, error)
}

type EmailAdapter interface {
	SendEmail(address, content string) (err error)
}
