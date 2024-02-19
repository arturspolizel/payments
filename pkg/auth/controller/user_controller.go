package controller

import (
	"fmt"

	"github.com/alexedwards/argon2id"
	"github.com/arturspolizel/payments/pkg/auth/interfaces"
	"github.com/arturspolizel/payments/pkg/auth/model"
	"github.com/arturspolizel/payments/utils"
)

type UserController struct {
	userRepository interfaces.UserRepository
	emailAdapter   interfaces.EmailAdapter
}

func NewUserController(paymentRepository interfaces.UserRepository, emailAdapter interfaces.EmailAdapter) *UserController {
	return &UserController{
		userRepository: paymentRepository,
		emailAdapter:   emailAdapter,
	}
}

func (c *UserController) Get(id uint) (model.User, error) {
	user, err := c.userRepository.Get(id)
	return user, err
}

func (c *UserController) Create(user model.User, password string) (uint, error) {
	// Validate user info
	user.Status = model.PendingActivation
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		// Log
		return 0, err
	}
	user.PasswordHash = hash
	id, err := c.userRepository.Create(user)
	if err != nil {
		// Check for existing user?
		return 0, err
	}

	validationEmail := model.ValidationEmail{
		UserId:    id,
		Code:      utils.RandStringBytes(6),
		Validated: false,
	}
	// Mocked email content
	validationContent := fmt.Sprintf("Welcome, %s! Your validation code is %s", user.Name, validationEmail.Code)
	err = c.emailAdapter.SendEmail(user.Email, validationContent)

	if err != nil {
		// Retry email sending? Cancel user creation?
		return id, err
	}

	return id, err
}

func (c UserController) Login(email string, password string) (model.User, error) {
	user, err := c.userRepository.GetByEmail(email)
	if err != nil {
		// Handle not found
		return user, err
	}

	match, err := argon2id.ComparePasswordAndHash(password, user.PasswordHash)
	if err != nil {
		// Handle not found
		return user, err
	}

	if !match {
		return user, err
	}

	// Create jwt token, return info
	return user, err
}

func (c UserController) Lock(id uint) error {
	panic("not implemented") // TODO: Implement
}
