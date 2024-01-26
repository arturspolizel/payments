package controller

import (
	"testing"

	"github.com/arturspolizel/payments/mocks"
	"github.com/arturspolizel/payments/model"
	"github.com/stretchr/testify/assert"
)

var mockedPayment = model.Payment{
	ID:         1,
	MerchantId: "1",
	Amount:     100,
	Tips:       100,
	Total:      200,
	Currency:   model.USD,
}

var paymentFromRequest = model.Payment{
	MerchantId: "1",
	Amount:     100,
	Tips:       100,
	Total:      200,
	Currency:   model.USD,
}

var paymentWithTotal = model.Payment{
	MerchantId: "1",
	Amount:     100,
	Tips:       100,
	Total:      200,
	Currency:   model.USD,
}

func TestGet(t *testing.T) {

	assert := assert.New(t)
	mockRepo := mocks.NewPaymentRepository(t)

	mockRepo.Mock.On("Get", uint(1)).Return(mockedPayment)
	defer mockRepo.AssertExpectations(t)

	paymentController := NewPaymentController(mockRepo)
	returnPayment := paymentController.Get(1)

	mockRepo.AssertExpectations(t)
	assert.Equal(mockedPayment, returnPayment)
}

func TestCreate(t *testing.T) {
	assert := assert.New(t)
	mockRepo := mocks.NewPaymentRepository(t)

	mockRepo.Mock.On("Create", paymentWithTotal).Return(uint(1))
	defer mockRepo.AssertExpectations(t)

	paymentController := NewPaymentController(mockRepo)
	id := paymentController.Create(paymentFromRequest)

	assert.Equal(uint(1), id)
}
