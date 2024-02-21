package controller

import (
	"errors"
	"fmt"

	"github.com/alexedwards/argon2id"
	"github.com/arturspolizel/payments/pkg/auth/interfaces"
	"github.com/arturspolizel/payments/pkg/auth/model"
	"github.com/arturspolizel/payments/utils"
)

type UserController struct {
	userRepository interfaces.UserRepository
	emailAdapter   interfaces.EmailAdapter
	jwtProcessor   utils.JwtProcessor
}

func NewUserController(paymentRepository interfaces.UserRepository, emailAdapter interfaces.EmailAdapter, jwtProcessor utils.JwtProcessor) *UserController {
	return &UserController{
		userRepository: paymentRepository,
		emailAdapter:   emailAdapter,
		jwtProcessor:   jwtProcessor,
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
	_, err = c.userRepository.CreateValidationEmail(validationEmail)
	if err != nil {
		// Retry email sending? Cancel user creation?
		return id, err
	}

	validationContent := fmt.Sprintf("Welcome, %s! Your validation code is %s", user.Name, validationEmail.Code)
	err = c.emailAdapter.SendEmail(user.Email, validationContent)

	if err != nil {
		// Retry email sending? Cancel user creation?
		return id, err
	}

	return id, err
}

func (c UserController) Login(email string, password string) (string, error) {
	user, err := c.userRepository.GetByEmail(email)
	if err != nil {
		// Handle not found
		var errNotFound *utils.ErrDatabaseNotFound
		if errors.As(err, &errNotFound) {
			return "", &model.ErrAuthenticationFailed{}
		}
		return "", err
	}

	if user.Status != model.Active {
		return "", &model.ErrInvalidUserStatus{
			Id: user.ID,
			AllowedStatuses: []model.UserStatus{
				model.Active,
			},
		}
	}

	match, err := argon2id.ComparePasswordAndHash(password, user.PasswordHash)
	if err != nil {
		return "", err
	}

	if !match {
		return "", &model.ErrAuthenticationFailed{}
	}

	// TODO: Create jwt token, return info
	token, err := c.jwtProcessor.NewToken(utils.TokenContext{
		Email:      user.Email,
		MerchantId: user.MerchantId,
	})

	return token, err
}

func (c UserController) Validate(code string) error {
	email, err := c.userRepository.GetEmailByCode(code)
	if err != nil {
		var errNotFound *utils.ErrDatabaseNotFound
		if errors.As(err, &errNotFound) {
			return &model.ErrInvalidEmailCode{}
		}
		return err
	}

	user, err := c.userRepository.Get(email.UserId)
	if err != nil {
		var errNotFound *utils.ErrDatabaseNotFound
		if errors.As(err, &errNotFound) {
			return &model.ErrInvalidEmailCode{}
		}
		return err
	}

	if email.Validated {
		// User already active
		return &model.ErrInvalidUserStatus{
			Id: user.ID,
			AllowedStatuses: []model.UserStatus{
				model.PendingActivation,
			},
		}
	}

	email.Validated = true
	user.Status = model.Active

	err = c.userRepository.Update(user)
	if err != nil {
		return err
	}

	err = c.userRepository.UpdateValidationEmail(email)

	return err

}
