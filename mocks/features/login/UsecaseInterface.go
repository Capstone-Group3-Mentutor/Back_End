// Code generated by mockery v2.14.1. DO NOT EDIT.

package mocks

import (
	login "be12/mentutor/features/login"

	mock "github.com/stretchr/testify/mock"
)

// UsecaseInterface is an autogenerated mock type for the UsecaseInterface type
type UsecaseInterface struct {
	mock.Mock
}

// Login provides a mock function with given fields: input
func (_m *UsecaseInterface) Login(input login.Core) (login.Core, string, error) {
	ret := _m.Called(input)

	var r0 login.Core
	if rf, ok := ret.Get(0).(func(login.Core) login.Core); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(login.Core)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(login.Core) string); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(login.Core) error); ok {
		r2 = rf(input)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type mockConstructorTestingTNewUsecaseInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewUsecaseInterface creates a new instance of UsecaseInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUsecaseInterface(t mockConstructorTestingTNewUsecaseInterface) *UsecaseInterface {
	mock := &UsecaseInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
