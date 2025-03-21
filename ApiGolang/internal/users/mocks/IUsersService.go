// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	context "context"

	users "github.com/EvansTrein/BlockbusterVHS/internal/users"
	mock "github.com/stretchr/testify/mock"
)

// IUsersService is an autogenerated mock type for the IUsersService type
type IUsersService struct {
	mock.Mock
}

// Register provides a mock function with given fields: ctx, data
func (_m *IUsersService) Register(ctx context.Context, data *users.RegisterRequest) (*users.ReqisterResponce, error) {
	ret := _m.Called(ctx, data)

	if len(ret) == 0 {
		panic("no return value specified for Register")
	}

	var r0 *users.ReqisterResponce
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *users.RegisterRequest) (*users.ReqisterResponce, error)); ok {
		return rf(ctx, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *users.RegisterRequest) *users.ReqisterResponce); ok {
		r0 = rf(ctx, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users.ReqisterResponce)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *users.RegisterRequest) error); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, data
func (_m *IUsersService) Update(ctx context.Context, data *users.UpdateRequest) (*users.UpdateResponce, error) {
	ret := _m.Called(ctx, data)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *users.UpdateResponce
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *users.UpdateRequest) (*users.UpdateResponce, error)); ok {
		return rf(ctx, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *users.UpdateRequest) *users.UpdateResponce); ok {
		r0 = rf(ctx, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users.UpdateResponce)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *users.UpdateRequest) error); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// User provides a mock function with given fields: ctx, id
func (_m *IUsersService) User(ctx context.Context, id uint) (*users.GetUserResponce, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for User")
	}

	var r0 *users.GetUserResponce
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) (*users.GetUserResponce, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint) *users.GetUserResponce); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users.GetUserResponce)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIUsersService creates a new instance of IUsersService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIUsersService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IUsersService {
	mock := &IUsersService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
