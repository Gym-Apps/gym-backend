// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/Gym-Apps/gym-backend/models"
	mock "github.com/stretchr/testify/mock"
)

// IUserRepository is an autogenerated mock type for the IUserRepository type
type IUserRepository struct {
	mock.Mock
}

// Login provides a mock function with given fields: ctx, phone
func (_m *IUserRepository) Login(ctx context.Context, phone string) (models.User, error) {
	ret := _m.Called(ctx, phone)

	var r0 models.User
	if rf, ok := ret.Get(0).(func(context.Context, string) models.User); ok {
		r0 = rf(ctx, phone)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, phone)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePassword provides a mock function with given fields: ctx, userID, password
func (_m *IUserRepository) UpdatePassword(ctx context.Context, userID uint, password string) error {
	ret := _m.Called(ctx, userID, password)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint, string) error); ok {
		r0 = rf(ctx, userID, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewIUserRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewIUserRepository creates a new instance of IUserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIUserRepository(t mockConstructorTestingTNewIUserRepository) *IUserRepository {
	mock := &IUserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
