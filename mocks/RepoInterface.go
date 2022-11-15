// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	mentor "be12/mentutor/features/mentor"

	mock "github.com/stretchr/testify/mock"
)

// RepoInterface is an autogenerated mock type for the RepoInterface type
type RepoInterface struct {
	mock.Mock
}

// AddScore provides a mock function with given fields: input
func (_m *RepoInterface) AddScore(input mentor.SubmissionCore) (mentor.SubmissionCore, error) {
	ret := _m.Called(input)

	var r0 mentor.SubmissionCore
	if rf, ok := ret.Get(0).(func(mentor.SubmissionCore) mentor.SubmissionCore); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(mentor.SubmissionCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(mentor.SubmissionCore) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTask provides a mock function with given fields: idTask, idClass
func (_m *RepoInterface) DeleteTask(idTask uint, idClass uint) (mentor.TaskCore, error) {
	ret := _m.Called(idTask, idClass)

	var r0 mentor.TaskCore
	if rf, ok := ret.Get(0).(func(uint, uint) mentor.TaskCore); ok {
		r0 = rf(idTask, idClass)
	} else {
		r0 = ret.Get(0).(mentor.TaskCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, uint) error); ok {
		r1 = rf(idTask, idClass)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EditProfileMentee provides a mock function with given fields: input
func (_m *RepoInterface) EditProfileMentee(input mentor.UserCore) (mentor.UserCore, error) {
	ret := _m.Called(input)

	var r0 mentor.UserCore
	if rf, ok := ret.Get(0).(func(mentor.UserCore) mentor.UserCore); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(mentor.UserCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(mentor.UserCore) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EditProfileMentor provides a mock function with given fields: input
func (_m *RepoInterface) EditProfileMentor(input mentor.UserCore) (mentor.UserCore, error) {
	ret := _m.Called(input)

	var r0 mentor.UserCore
	if rf, ok := ret.Get(0).(func(mentor.UserCore) mentor.UserCore); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(mentor.UserCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(mentor.UserCore) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EditTask provides a mock function with given fields: input
func (_m *RepoInterface) EditTask(input mentor.TaskCore) (mentor.TaskCore, error) {
	ret := _m.Called(input)

	var r0 mentor.TaskCore
	if rf, ok := ret.Get(0).(func(mentor.TaskCore) mentor.TaskCore); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(mentor.TaskCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(mentor.TaskCore) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllTask provides a mock function with given fields:
func (_m *RepoInterface) GetAllTask() ([]mentor.TaskCore, error) {
	ret := _m.Called()

	var r0 []mentor.TaskCore
	if rf, ok := ret.Get(0).(func() []mentor.TaskCore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]mentor.TaskCore)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSingleMentee provides a mock function with given fields: id
func (_m *RepoInterface) GetSingleMentee(id uint) (mentor.UserCore, error) {
	ret := _m.Called(id)

	var r0 mentor.UserCore
	if rf, ok := ret.Get(0).(func(uint) mentor.UserCore); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(mentor.UserCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSingleMentor provides a mock function with given fields: id
func (_m *RepoInterface) GetSingleMentor(id uint) (mentor.UserCore, error) {
	ret := _m.Called(id)

	var r0 mentor.UserCore
	if rf, ok := ret.Get(0).(func(uint) mentor.UserCore); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(mentor.UserCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSingleTask provides a mock function with given fields: id
func (_m *RepoInterface) GetSingleTask(id uint) (mentor.TaskCore, error) {
	ret := _m.Called(id)

	var r0 mentor.TaskCore
	if rf, ok := ret.Get(0).(func(uint) mentor.TaskCore); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(mentor.TaskCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSubmission provides a mock function with given fields: id
func (_m *RepoInterface) GetSubmission(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTaskSub provides a mock function with given fields: id
func (_m *RepoInterface) GetTaskSub(id uint) (mentor.TaskCore, []mentor.SubmissionCore, error) {
	ret := _m.Called(id)

	var r0 mentor.TaskCore
	if rf, ok := ret.Get(0).(func(uint) mentor.TaskCore); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(mentor.TaskCore)
	}

	var r1 []mentor.SubmissionCore
	if rf, ok := ret.Get(1).(func(uint) []mentor.SubmissionCore); ok {
		r1 = rf(id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]mentor.SubmissionCore)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(uint) error); ok {
		r2 = rf(id)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// InsertTask provides a mock function with given fields: input
func (_m *RepoInterface) InsertTask(input mentor.TaskCore) (mentor.TaskCore, error) {
	ret := _m.Called(input)

	var r0 mentor.TaskCore
	if rf, ok := ret.Get(0).(func(mentor.TaskCore) mentor.TaskCore); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(mentor.TaskCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(mentor.TaskCore) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRepoInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepoInterface creates a new instance of RepoInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepoInterface(t mockConstructorTestingTNewRepoInterface) *RepoInterface {
	mock := &RepoInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}