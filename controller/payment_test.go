package controller

import (
	"testing"

	"github.com/arturspolizel/payments/mocks"
	"github.com/arturspolizel/payments/model"
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
