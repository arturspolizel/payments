package controller

import (
	"testing"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/arturspolizel/payments/mocks"
	"github.com/arturspolizel/payments/pkg/auth/model"
	"github.com/arturspolizel/payments/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/mock"
)

type userDeps struct {
	userRepository mocks.UserRepository
	emailAdapter   mocks.EmailAdapter
	jwtProcessor   mocks.JwtProcessor
}

var testPassword = "test"
var passwordHash, _ = argon2id.CreateHash(testPassword, argon2id.DefaultParams)

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
				password: testPassword,
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

			if got, err := controller.Create(tt.args.user, tt.args.password); got != tt.want && !utils.CheckTestError(err, tt.wantErr) {
				t.Errorf("MerchantController.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserController_Login(t *testing.T) {
	type args struct {
		email    string
		password string
	}
	tests := []struct {
		name    string
		c       UserController
		args    args
		want    string
		wantErr error

		on     func(*userDeps)
		assert func(*userDeps)
	}{
		{
			name: "Success",
			args: args{
				email:    "test@test.com",
				password: testPassword,
			},
			want:    "mockToken",
			wantErr: nil,
			on: func(ud *userDeps) {
				ud.userRepository.Mock.On("GetByEmail", "test@test.com").Return(model.User{
					ID:           1,
					Name:         "test",
					MerchantId:   1,
					Email:        "test@test.com",
					Status:       model.Active,
					PasswordHash: passwordHash,
				}, nil)
				ud.jwtProcessor.Mock.On("NewToken", utils.TokenContext{
					Email:      "test@test.com",
					MerchantId: 1,
					RegisteredClaims: jwt.RegisteredClaims{
						Issuer:    "http://localhost/payments/auth", //TODO: extract to correct domain
						Subject:   "test@test.com",
						ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
					},
				}).Return("mockToken", nil)
			},
			assert: func(ud *userDeps) {
				ud.emailAdapter.AssertExpectations(t)
				ud.jwtProcessor.AssertExpectations(t)
				ud.userRepository.AssertExpectations(t)
			},
		},
		{
			name: "Wrong password",
			args: args{
				email:    "test@test.com",
				password: "wrong",
			},
			want:    "",
			wantErr: &model.ErrAuthenticationFailed{},
			on: func(ud *userDeps) {
				ud.userRepository.Mock.On("GetByEmail", "test@test.com").Return(model.User{
					ID:           1,
					Name:         "test",
					MerchantId:   1,
					Email:        "test@test.com",
					Status:       model.Active,
					PasswordHash: passwordHash,
				}, nil)
			},
			assert: func(ud *userDeps) {
				ud.emailAdapter.AssertExpectations(t)
				ud.jwtProcessor.AssertExpectations(t)
				ud.userRepository.AssertExpectations(t)
			},
		},
		{
			name: "User not active",
			args: args{
				email:    "test@test.com",
				password: testPassword,
			},
			want:    "",
			wantErr: &model.ErrInvalidUserStatus{},
			on: func(ud *userDeps) {
				ud.userRepository.Mock.On("GetByEmail", "test@test.com").Return(model.User{
					ID:           1,
					Name:         "test",
					MerchantId:   1,
					Email:        "test@test.com",
					Status:       model.PendingActivation,
					PasswordHash: passwordHash,
				}, nil)
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

			got, err := controller.Login(tt.args.email, tt.args.password)
			if (err != nil) && !utils.CheckTestError(err, tt.wantErr) {
				t.Errorf("UserController.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UserController.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserController_Validate(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name    string
		c       UserController
		args    args
		wantErr error

		on     func(*userDeps)
		assert func(*userDeps)
	}{
		{
			name: "Success",
			args: args{
				code: "mockCode",
			},
			wantErr: nil,
			on: func(ud *userDeps) {
				ud.userRepository.Mock.On("GetEmailByCode", "mockCode").Return(model.ValidationEmail{
					ID:        1,
					UserId:    1,
					Code:      "mockCode",
					Validated: false,
				}, nil)
				ud.userRepository.Mock.On("Get", uint(1)).Return(model.User{
					ID:         1,
					Name:       "test",
					MerchantId: 1,
					Email:      "test@test.com",
					Status:     model.PendingActivation,
				}, nil)
				ud.userRepository.Mock.On("Update", model.User{
					ID:         1,
					Name:       "test",
					MerchantId: 1,
					Email:      "test@test.com",
					Status:     model.Active,
				}).Return(nil)
				ud.userRepository.Mock.On("UpdateValidationEmail", model.ValidationEmail{
					ID:        1,
					UserId:    1,
					Code:      "mockCode",
					Validated: true,
				}).Return(nil)
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

			if err := controller.Validate(tt.args.code); !utils.CheckTestError(err, tt.wantErr) {
				t.Errorf("UserController.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
