package controller

import (
	"errors"
	"testing"

	"github.com/alexedwards/argon2id"
	"github.com/arturspolizel/payments/mocks"
	"github.com/arturspolizel/payments/pkg/auth/model"
	"github.com/stretchr/testify/mock"
)

type userDeps struct {
	userRepository mocks.UserRepository
	emailAdapter   mocks.EmailAdapter
	jwtProcessor   mocks.JwtProcessor
}

func TestUserController_Create(t *testing.T) {
	type args struct {
		user     model.User
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    uint
		wantErr error

		on     func(*userDeps)
		assert func(*userDeps)
	}{
		{
			name: "Success",
			args: args{
				user: model.User{
					Name:       "test",
					MerchantId: 1,
					Email:      "test@test.com",
				},
				password: "test",
			},
			want:    1,
			wantErr: nil,
			on: func(ud *userDeps) {
				ud.userRepository.Mock.On("Create", mock.MatchedBy(func(u model.User) bool {
					validHash, _ := argon2id.ComparePasswordAndHash("test", u.PasswordHash)
					return u.MerchantId == 1 && u.Name == "test" && u.Email == "test@test.com" && u.Status == model.PendingActivation && validHash
				})).Return(uint(1), nil)

				ud.userRepository.Mock.On("CreateValidationEmail", mock.MatchedBy(func(e model.ValidationEmail) bool {
					return e.UserId == 1 && !e.Validated
				})).Return(uint(1), nil)

				ud.emailAdapter.Mock.On("SendEmail", "test@test.com", mock.Anything).Return(nil)
			},
			assert: func(ud *userDeps) {
				ud.emailAdapter.AssertExpectations(t)
				ud.jwtProcessor.AssertExpectations(t)
				ud.userRepository.AssertExpectations(t)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			depFields := userDeps{
				userRepository: *mocks.NewUserRepository(t),
				emailAdapter:   *mocks.NewEmailAdapter(t),
				jwtProcessor:   *mocks.NewJwtProcessor(t),
			}
			tt.on(&depFields)
			defer tt.assert(&depFields)

			controller := NewUserController(&depFields.userRepository, &depFields.emailAdapter, &depFields.jwtProcessor)

			if got, err := controller.Create(tt.args.user, tt.args.password); got != tt.want && errors.Is(err, tt.wantErr) {
				t.Errorf("MerchantController.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
