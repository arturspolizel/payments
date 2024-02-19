package controller

import (
	"errors"
	"testing"

	"github.com/arturspolizel/payments/mocks"
	"github.com/arturspolizel/payments/pkg/payment/model"
	"github.com/stretchr/testify/assert"
)

var mockedPayment = model.Payment{
	ID:         1,
	MerchantId: uint(1),
	Amount:     100,
	Tips:       100,
	Total:      200,
	Currency:   model.USD,
}

var authorizedPayment = model.Payment{
	ID:         1,
	MerchantId: uint(1),
	Amount:     100,
	Tips:       100,
	Total:      200,
	Currency:   model.USD,
	Status:     model.Authorized,
}

var paymentFromRequest = model.Payment{
	MerchantId: uint(1),
	Amount:     100,
	Tips:       100,
	Total:      200,
	Currency:   model.USD,
}

var paymentWithTotal = model.Payment{
	MerchantId: uint(1),
	Amount:     100,
	Tips:       100,
	Total:      200,
	Currency:   model.USD,
	Status:     model.Captured,
}

type paymentDeps struct {
	paymentRepo mocks.PaymentRepository
}

func TestGet(t *testing.T) {

	assert := assert.New(t)
	mockRepo := mocks.NewPaymentRepository(t)

	mockRepo.Mock.On("Get", uint(1)).Return(mockedPayment, nil)
	defer mockRepo.AssertExpectations(t)

	paymentController := NewPaymentController(mockRepo)
	returnPayment, err := paymentController.Get(1)

	mockRepo.AssertExpectations(t)
	assert.Equal(mockedPayment, returnPayment)
	assert.Empty(err)
}

func TestCreate(t *testing.T) {
	assert := assert.New(t)
	mockRepo := mocks.NewPaymentRepository(t)

	mockRepo.Mock.On("Create", paymentWithTotal).Return(uint(1), nil)
	defer mockRepo.AssertExpectations(t)

	paymentController := NewPaymentController(mockRepo)
	id, err := paymentController.Create(paymentFromRequest)

	assert.Equal(uint(1), id)
	assert.Empty(err)
}

func TestPaymentController_Authorize(t *testing.T) {
	type args struct {
		payment model.Payment
	}
	tests := []struct {
		name    string
		args    args
		want    uint
		wantErr error

		on     func(*paymentDeps)
		assert func(*paymentDeps)
	}{
		{
			name: "Success",
			args: args{
				payment: model.Payment{
					MerchantId: uint(1),
					Amount:     100,
					Tips:       100,
					Currency:   model.USD,
				},
			},
			want:    1,
			wantErr: nil,
			on: func(pd *paymentDeps) {
				pd.paymentRepo.Mock.On("Create", model.Payment{
					MerchantId: uint(1),
					Amount:     100,
					Tips:       100,
					Currency:   model.USD,
					Total:      200,
					Status:     model.Authorized,
				}).Return(uint(1), nil)
			},
			assert: func(pd *paymentDeps) {
				pd.paymentRepo.AssertExpectations(t)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			depFields := paymentDeps{
				paymentRepo: *mocks.NewPaymentRepository(t),
			}

			tt.on(&depFields)
			defer tt.assert(&depFields)

			controller := NewPaymentController(&depFields.paymentRepo)
			got, err := controller.Authorize(tt.args.payment)
			if err != nil && errors.Is(err, tt.wantErr) {
				t.Errorf("PaymentController.Authorize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PaymentController.Authorize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPaymentController_Capture(t *testing.T) {
	type args struct {
		id     uint
		amount int
		tips   int
	}
	tests := []struct {
		name    string
		args    args
		want    uint
		wantErr error

		on     func(*paymentDeps)
		assert func(*paymentDeps)
	}{
		{
			name: "Success",
			args: args{
				id:     1,
				amount: 200,
				tips:   200,
			},
			want:    1,
			wantErr: nil,
			on: func(pd *paymentDeps) {
				pd.paymentRepo.Mock.On("Get", uint(1)).
					Return(authorizedPayment, nil)
				pd.paymentRepo.Mock.On("Update", model.Payment{
					ID:         1,
					MerchantId: uint(1),
					Amount:     200,
					Tips:       200,
					Currency:   model.USD,
					Total:      400,
					Status:     model.Captured,
				}).Return(nil)
			},
			assert: func(pd *paymentDeps) {
				pd.paymentRepo.AssertExpectations(t)
			},
		},
		{
			name: "Invalid status",
			args: args{
				id:     1,
				amount: 200,
				tips:   200,
			},
			want:    1,
			wantErr: &model.ErrInvalidPaymentStatus{},
			on: func(pd *paymentDeps) {
				pd.paymentRepo.Mock.On("Get", uint(1)).
					Return(model.Payment{
						ID:         1,
						MerchantId: uint(1),
						Amount:     200,
						Tips:       200,
						Currency:   model.USD,
						Total:      400,
						Status:     model.Captured,
					}, nil)
			},
			assert: func(pd *paymentDeps) {
				pd.paymentRepo.AssertExpectations(t)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			depFields := paymentDeps{
				paymentRepo: *mocks.NewPaymentRepository(t),
			}

			tt.on(&depFields)
			defer tt.assert(&depFields)

			controller := NewPaymentController(&depFields.paymentRepo)
			err := controller.Capture(tt.args.id, tt.args.amount, tt.args.tips)
			if err != nil && errors.Is(err, tt.wantErr) {
				t.Errorf("PaymentController.Capture() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestPaymentController_Refund(t *testing.T) {
	type args struct {
		id     uint
		amount int
		tips   int
	}
	tests := []struct {
		name    string
		args    args
		want    uint
		wantErr error

		on     func(*paymentDeps)
		assert func(*paymentDeps)
	}{
		{
			name: "Success",
			args: args{
				id:     1,
				amount: 100,
				tips:   100,
			},
			want:    1,
			wantErr: nil,
			on: func(pd *paymentDeps) {
				pd.paymentRepo.Mock.On("Get", uint(1)).
					Return(model.Payment{
						ID:         1,
						MerchantId: uint(1),
						Amount:     100,
						Tips:       100,
						Currency:   model.USD,
						Total:      200,
						Status:     model.Captured,
					}, nil)
				pd.paymentRepo.Mock.On("CreateRefund", model.Refund{
					PaymentId: 1,
					Amount:    100,
					Tips:      100,
					Total:     200,
				}).Return(uint(1), nil)
			},
			assert: func(pd *paymentDeps) {
				pd.paymentRepo.AssertExpectations(t)
			},
		},
		{
			name: "Invalid status",
			args: args{
				id:     1,
				amount: 200,
				tips:   200,
			},
			want:    1,
			wantErr: &model.ErrInvalidPaymentStatus{},
			on: func(pd *paymentDeps) {
				pd.paymentRepo.Mock.On("Get", uint(1)).
					Return(model.Payment{
						ID:         1,
						MerchantId: uint(1),
						Amount:     200,
						Tips:       200,
						Currency:   model.USD,
						Total:      400,
						Status:     model.Authorized,
					}, nil)
			},
			assert: func(pd *paymentDeps) {
				pd.paymentRepo.AssertExpectations(t)
			},
		},
		{
			name: "Invalid amount",
			args: args{
				id:     1,
				amount: 200,
				tips:   200,
			},
			want:    1,
			wantErr: &model.ErrInvalidPaymentStatus{},
			on: func(pd *paymentDeps) {
				pd.paymentRepo.Mock.On("Get", uint(1)).
					Return(model.Payment{
						ID:         1,
						MerchantId: uint(1),
						Amount:     100,
						Tips:       100,
						Currency:   model.USD,
						Total:      400,
						Status:     model.Captured,
					}, nil)
			},
			assert: func(pd *paymentDeps) {
				pd.paymentRepo.AssertExpectations(t)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			depFields := paymentDeps{
				paymentRepo: *mocks.NewPaymentRepository(t),
			}

			tt.on(&depFields)
			defer tt.assert(&depFields)

			controller := NewPaymentController(&depFields.paymentRepo)
			err := controller.Refund(tt.args.id, tt.args.amount, tt.args.tips)
			if err != nil && errors.Is(err, tt.wantErr) {
				t.Errorf("PaymentController.Capture() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestPaymentController_Void(t *testing.T) {
	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		args    args
		want    uint
		wantErr error

		on     func(*paymentDeps)
		assert func(*paymentDeps)
	}{
		{
			name: "Success",
			args: args{
				id: 1,
			},
			want:    1,
			wantErr: nil,
			on: func(pd *paymentDeps) {
				pd.paymentRepo.Mock.On("Get", uint(1)).
					Return(authorizedPayment, nil)
				pd.paymentRepo.Mock.On("Update", model.Payment{
					ID:         1,
					MerchantId: uint(1),
					Amount:     100,
					Tips:       100,
					Currency:   model.USD,
					Total:      200,
					Status:     model.Void,
				}).Return(nil)
			},
			assert: func(pd *paymentDeps) {
				pd.paymentRepo.AssertExpectations(t)
			},
		},
		{
			name: "Invalid status",
			args: args{
				id: 1,
			},
			want:    1,
			wantErr: &model.ErrInvalidPaymentStatus{},
			on: func(pd *paymentDeps) {
				pd.paymentRepo.Mock.On("Get", uint(1)).
					Return(model.Payment{
						ID:         1,
						MerchantId: uint(1),
						Amount:     200,
						Tips:       200,
						Currency:   model.USD,
						Total:      400,
						Status:     model.Captured,
					}, nil)
			},
			assert: func(pd *paymentDeps) {
				pd.paymentRepo.AssertExpectations(t)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			depFields := paymentDeps{
				paymentRepo: *mocks.NewPaymentRepository(t),
			}

			tt.on(&depFields)
			defer tt.assert(&depFields)

			controller := NewPaymentController(&depFields.paymentRepo)
			err := controller.Void(tt.args.id)
			if err != nil && errors.Is(err, tt.wantErr) {
				t.Errorf("PaymentController.Capture() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
