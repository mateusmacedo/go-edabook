// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	contract "goedabook/internal/acl/user/application/contract"
)

// Handler is an autogenerated mock type for the Handler type
type Handler struct {
	mock.Mock
}

// Handle provides a mock function with given fields: ctx, cmd
func (_m *Handler) Handle(ctx context.Context, cmd contract.Command) (contract.HandlerResult, contract.HandlerErr) {
	ret := _m.Called(ctx, cmd)

	if len(ret) == 0 {
		panic("no return value specified for Handle")
	}

	var r0 contract.HandlerResult
	var r1 contract.HandlerErr
	if rf, ok := ret.Get(0).(func(context.Context, contract.Command) (contract.HandlerResult, contract.HandlerErr)); ok {
		return rf(ctx, cmd)
	}
	if rf, ok := ret.Get(0).(func(context.Context, contract.Command) contract.HandlerResult); ok {
		r0 = rf(ctx, cmd)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(contract.HandlerResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, contract.Command) contract.HandlerErr); ok {
		r1 = rf(ctx, cmd)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(contract.HandlerErr)
		}
	}

	return r0, r1
}

// NewHandler creates a new instance of Handler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *Handler {
	mock := &Handler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}