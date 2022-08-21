// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	request "github.com/Gym-Apps/gym-backend/dto/request"
	mock "github.com/stretchr/testify/mock"

	response "github.com/Gym-Apps/gym-backend/dto/response"
)

// IUserService is an autogenerated mock type for the IUserService type
type IUserService struct {
	mock.Mock
}

// Login provides a mock function with given fields: userLoginRequest
func (_m *IUserService) Login(userLoginRequest request.UserLoginDTO) (response.UserLoginDTO, error) {
	ret := _m.Called(userLoginRequest)

	var r0 response.UserLoginDTO
	if rf, ok := ret.Get(0).(func(request.UserLoginDTO) response.UserLoginDTO); ok {
		r0 = rf(userLoginRequest)
	} else {
		r0 = ret.Get(0).(response.UserLoginDTO)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(request.UserLoginDTO) error); ok {
		r1 = rf(userLoginRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIUserService interface {
	mock.TestingT
	Cleanup(func())
}

// NewIUserService creates a new instance of IUserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIUserService(t mockConstructorTestingTNewIUserService) *IUserService {
	mock := &IUserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
