// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	time "time"

	model "github.com/arturspolizel/payments/pkg/payment/model"
	mock "github.com/stretchr/testify/mock"
)

// PaymentRepository is an autogenerated mock type for the PaymentRepository type
type PaymentRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *PaymentRepository) Create(_a0 model.Payment) (uint, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 uint
	var r1 error
	if rf, ok := ret.Get(0).(func(model.Payment) (uint, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(model.Payment) uint); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(uint)
	}

	if rf, ok := ret.Get(1).(func(model.Payment) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRefund provides a mock function with given fields: refund
func (_m *PaymentRepository) CreateRefund(refund model.Refund) (uint, error) {
	ret := _m.Called(refund)

	if len(ret) == 0 {
		panic("no return value specified for CreateRefund")
	}

	var r0 uint
	var r1 error
	if rf, ok := ret.Get(0).(func(model.Refund) (uint, error)); ok {
		return rf(refund)
	}
	if rf, ok := ret.Get(0).(func(model.Refund) uint); ok {
		r0 = rf(refund)
	} else {
		r0 = ret.Get(0).(uint)
	}

	if rf, ok := ret.Get(1).(func(model.Refund) error); ok {
		r1 = rf(refund)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: _a0
func (_m *PaymentRepository) Get(_a0 uint) (model.Payment, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 model.Payment
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (model.Payment, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(uint) model.Payment); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(model.Payment)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: startId, pageSize, startDate, endDate
func (_m *PaymentRepository) List(startId uint, pageSize uint, startDate time.Time, endDate time.Time) ([]model.Payment, error) {
	ret := _m.Called(startId, pageSize, startDate, endDate)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []model.Payment
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, uint, time.Time, time.Time) ([]model.Payment, error)); ok {
		return rf(startId, pageSize, startDate, endDate)
	}
	if rf, ok := ret.Get(0).(func(uint, uint, time.Time, time.Time) []model.Payment); ok {
		r0 = rf(startId, pageSize, startDate, endDate)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Payment)
		}
	}

	if rf, ok := ret.Get(1).(func(uint, uint, time.Time, time.Time) error); ok {
		r1 = rf(startId, pageSize, startDate, endDate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0
func (_m *PaymentRepository) Update(_a0 model.Payment) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(model.Payment) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewPaymentRepository creates a new instance of PaymentRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPaymentRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *PaymentRepository {
	mock := &PaymentRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
