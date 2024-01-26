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

func TestGet(t *testing.T) {
	assert := assert.New(t)
	mockRepo := mocks.NewPaymentRepository(t)

	mockRepo.Mock.On("Get", uint(1)).Return(mockedPayment)

	paymentController := NewPaymentController(mockRepo)
	returnPayment := paymentController.Get(1)

	mockRepo.AssertExpectations(t)
	assert.Equal(returnPayment, mockedPayment)
}

func TestCreate(t *testing.T) {

}
