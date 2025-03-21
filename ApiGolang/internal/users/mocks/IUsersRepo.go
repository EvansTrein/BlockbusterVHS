// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	context "context"

	users "github.com/EvansTrein/BlockbusterVHS/internal/users"
	mock "github.com/stretchr/testify/mock"
)

// IUsersRepo is an autogenerated mock type for the IUsersRepo type
type IUsersRepo struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, data
func (_m *IUsersRepo) Create(ctx context.Context, data *users.RegisterRequest) (uint, error) {
	ret := _m.Called(ctx, data)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 uint
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *users.RegisterRequest) (uint, error)); ok {
		return rf(ctx, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *users.RegisterRequest) uint); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Get(0).(uint)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *users.RegisterRequest) error); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExistsByID provides a mock function with given fields: ctx, id
func (_m *IUsersRepo) ExistsByID(ctx context.Context, id uint) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for ExistsByID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUserData provides a mock function with given fields: ctx, id, data
func (_m *IUsersRepo) GetUserData(ctx context.Context, id uint, data *users.GetUserResponce) error {
	ret := _m.Called(ctx, id, data)

	if len(ret) == 0 {
		panic("no return value specified for GetUserData")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint, *users.GetUserResponce) error); ok {
		r0 = rf(ctx, id, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, data
func (_m *IUsersRepo) Update(ctx context.Context, data *users.UpdateRequest) error {
	ret := _m.Called(ctx, data)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *users.UpdateRequest) error); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIUsersRepo creates a new instance of IUsersRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIUsersRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *IUsersRepo {
	mock := &IUsersRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
