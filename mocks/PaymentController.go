// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	model "github.com/arturspolizel/payments/model"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// PaymentController is an autogenerated mock type for the PaymentController type
type PaymentController struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *PaymentController) Create(_a0 model.Payment) (uint, error) {
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

// Get provides a mock function with given fields: _a0
func (_m *PaymentController) Get(_a0 uint) (model.Payment, error) {
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

// List provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *PaymentController) List(_a0 uint, _a1 uint, _a2 time.Time, _a3 time.Time) ([]model.Payment, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []model.Payment
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, uint, time.Time, time.Time) ([]model.Payment, error)); ok {
		return rf(_a0, _a1, _a2, _a3)
	}
	if rf, ok := ret.Get(0).(func(uint, uint, time.Time, time.Time) []model.Payment); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Payment)
		}
	}

	if rf, ok := ret.Get(1).(func(uint, uint, time.Time, time.Time) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPaymentController creates a new instance of PaymentController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPaymentController(t interface {
	mock.TestingT
	Cleanup(func())
}) *PaymentController {
	mock := &PaymentController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
