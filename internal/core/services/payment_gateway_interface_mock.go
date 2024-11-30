// Code generated by mockery v2.49.0. DO NOT EDIT.

package services

import (
	dtos "tech-challenge-fase-1/internal/core/dtos"

	mock "github.com/stretchr/testify/mock"
)

// MockPaymentGatewayInterface is an autogenerated mock type for the PaymentGatewayInterface type
type MockPaymentGatewayInterface struct {
	mock.Mock
}

type MockPaymentGatewayInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPaymentGatewayInterface) EXPECT() *MockPaymentGatewayInterface_Expecter {
	return &MockPaymentGatewayInterface_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: order, method
func (_m *MockPaymentGatewayInterface) Execute(order *dtos.OrderDTO, method dtos.MethodType) (*dtos.CheckoutDTO, error) {
	ret := _m.Called(order, method)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 *dtos.CheckoutDTO
	var r1 error
	if rf, ok := ret.Get(0).(func(*dtos.OrderDTO, dtos.MethodType) (*dtos.CheckoutDTO, error)); ok {
		return rf(order, method)
	}
	if rf, ok := ret.Get(0).(func(*dtos.OrderDTO, dtos.MethodType) *dtos.CheckoutDTO); ok {
		r0 = rf(order, method)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dtos.CheckoutDTO)
		}
	}

	if rf, ok := ret.Get(1).(func(*dtos.OrderDTO, dtos.MethodType) error); ok {
		r1 = rf(order, method)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPaymentGatewayInterface_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockPaymentGatewayInterface_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - order *dtos.OrderDTO
//   - method dtos.MethodType
func (_e *MockPaymentGatewayInterface_Expecter) Execute(order interface{}, method interface{}) *MockPaymentGatewayInterface_Execute_Call {
	return &MockPaymentGatewayInterface_Execute_Call{Call: _e.mock.On("Execute", order, method)}
}

func (_c *MockPaymentGatewayInterface_Execute_Call) Run(run func(order *dtos.OrderDTO, method dtos.MethodType)) *MockPaymentGatewayInterface_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*dtos.OrderDTO), args[1].(dtos.MethodType))
	})
	return _c
}

func (_c *MockPaymentGatewayInterface_Execute_Call) Return(_a0 *dtos.CheckoutDTO, _a1 error) *MockPaymentGatewayInterface_Execute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPaymentGatewayInterface_Execute_Call) RunAndReturn(run func(*dtos.OrderDTO, dtos.MethodType) (*dtos.CheckoutDTO, error)) *MockPaymentGatewayInterface_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPaymentGatewayInterface creates a new instance of MockPaymentGatewayInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPaymentGatewayInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPaymentGatewayInterface {
	mock := &MockPaymentGatewayInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}