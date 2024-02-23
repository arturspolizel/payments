// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	model "github.com/arturspolizel/payments/pkg/auth/model"
	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *UserRepository) Create(_a0 model.User) (uint, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 uint
	var r1 error
	if rf, ok := ret.Get(0).(func(model.User) (uint, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(model.User) uint); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(uint)
	}

	if rf, ok := ret.Get(1).(func(model.User) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateValidationEmail provides a mock function with given fields: _a0
func (_m *UserRepository) CreateValidationEmail(_a0 model.ValidationEmail) (uint, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for CreateValidationEmail")
	}

	var r0 uint
	var r1 error
	if rf, ok := ret.Get(0).(func(model.ValidationEmail) (uint, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(model.ValidationEmail) uint); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(uint)
	}

	if rf, ok := ret.Get(1).(func(model.ValidationEmail) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: _a0
func (_m *UserRepository) Get(_a0 uint) (model.User, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (model.User, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(uint) model.User); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(model.User)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByEmail provides a mock function with given fields: email
func (_m *UserRepository) GetByEmail(email string) (model.User, error) {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for GetByEmail")
	}

	var r0 model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (model.User, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) model.User); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(model.User)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEmailByCode provides a mock function with given fields: code
func (_m *UserRepository) GetEmailByCode(code string) (model.ValidationEmail, error) {
	ret := _m.Called(code)

	if len(ret) == 0 {
		panic("no return value specified for GetEmailByCode")
	}

	var r0 model.ValidationEmail
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (model.ValidationEmail, error)); ok {
		return rf(code)
	}
	if rf, ok := ret.Get(0).(func(string) model.ValidationEmail); ok {
		r0 = rf(code)
	} else {
		r0 = ret.Get(0).(model.ValidationEmail)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(code)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0
func (_m *UserRepository) Update(_a0 model.User) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(model.User) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateValidationEmail provides a mock function with given fields: _a0
func (_m *UserRepository) UpdateValidationEmail(_a0 model.ValidationEmail) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for UpdateValidationEmail")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(model.ValidationEmail) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
