// Code generated by mockery v2.14.1. DO NOT EDIT.

package mocks

import (
	mentee "be12/mentutor/features/mentee"

	mock "github.com/stretchr/testify/mock"
)

// UseCaseInterface is an autogenerated mock type for the UseCaseInterface type
type UseCaseInterface struct {
	mock.Mock
}

// GetAll provides a mock function with given fields:
func (_m *UseCaseInterface) GetAll() ([]mentee.Status, []mentee.CommentsCore, []mentee.CommentsCore, error) {
	ret := _m.Called()

	var r0 []mentee.Status
	if rf, ok := ret.Get(0).(func() []mentee.Status); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]mentee.Status)
		}
	}

	var r1 []mentee.CommentsCore
	if rf, ok := ret.Get(1).(func() []mentee.CommentsCore); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]mentee.CommentsCore)
		}
	}

	var r2 []mentee.CommentsCore
	if rf, ok := ret.Get(2).(func() []mentee.CommentsCore); ok {
		r2 = rf()
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).([]mentee.CommentsCore)
		}
	}

	var r3 error
	if rf, ok := ret.Get(3).(func() error); ok {
		r3 = rf()
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// GetTask provides a mock function with given fields: idClass, role
func (_m *UseCaseInterface) GetTask(idClass uint, role string) ([]mentee.Task, error) {
	ret := _m.Called(idClass, role)

	var r0 []mentee.Task
	if rf, ok := ret.Get(0).(func(uint, string) []mentee.Task); ok {
		r0 = rf(idClass, role)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]mentee.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, string) error); ok {
		r1 = rf(idClass, role)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: data
func (_m *UseCaseInterface) Insert(data mentee.CommentsCore) (mentee.CommentsCore, error) {
	ret := _m.Called(data)

	var r0 mentee.CommentsCore
	if rf, ok := ret.Get(0).(func(mentee.CommentsCore) mentee.CommentsCore); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(mentee.CommentsCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(mentee.CommentsCore) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertStatus provides a mock function with given fields: data, token
func (_m *UseCaseInterface) InsertStatus(data mentee.Status, token int) (mentee.Status, error) {
	ret := _m.Called(data, token)

	var r0 mentee.Status
	if rf, ok := ret.Get(0).(func(mentee.Status, int) mentee.Status); ok {
		r0 = rf(data, token)
	} else {
		r0 = ret.Get(0).(mentee.Status)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(mentee.Status, int) error); ok {
		r1 = rf(data, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertSub provides a mock function with given fields: data
func (_m *UseCaseInterface) InsertSub(data mentee.Submission) (mentee.Submission, error) {
	ret := _m.Called(data)

	var r0 mentee.Submission
	if rf, ok := ret.Get(0).(func(mentee.Submission) mentee.Submission); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(mentee.Submission)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(mentee.Submission) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUseCaseInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewUseCaseInterface creates a new instance of UseCaseInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUseCaseInterface(t mockConstructorTestingTNewUseCaseInterface) *UseCaseInterface {
	mock := &UseCaseInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
