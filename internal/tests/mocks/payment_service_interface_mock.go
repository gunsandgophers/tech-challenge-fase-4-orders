package mocks

import (
	dtos "tech-challenge-fase-1/internal/core/dtos"

	mock "github.com/stretchr/testify/mock"
)

// MockPaymentServiceInterface is an autogenerated mock type for the PaymentServiceInterface type
type MockPaymentServiceInterface struct {
	mock.Mock
}

type MockPaymentServiceInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPaymentServiceInterface) EXPECT() *MockPaymentServiceInterface_Expecter {
	return &MockPaymentServiceInterface_Expecter{mock: &_m.Mock}
}

// CreatePayment provides a mock function with given fields: orderID, amount
func (_m *MockPaymentServiceInterface) CreatePayment(orderID string, amount float64) (*dtos.CheckoutDTO, error) {
	ret := _m.Called(orderID, amount)

	if len(ret) == 0 {
		panic("no return value specified for CreatePayment")
	}

	var r0 *dtos.CheckoutDTO
	var r1 error
	if rf, ok := ret.Get(0).(func(string, float64) (*dtos.CheckoutDTO, error)); ok {
		return rf(orderID, amount)
	}
	if rf, ok := ret.Get(0).(func(string, float64) *dtos.CheckoutDTO); ok {
		r0 = rf(orderID, amount)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dtos.CheckoutDTO)
		}
	}

	if rf, ok := ret.Get(1).(func(string, float64) error); ok {
		r1 = rf(orderID, amount)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPaymentServiceInterface_CreatePayment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreatePayment'
type MockPaymentServiceInterface_CreatePayment_Call struct {
	*mock.Call
}

// CreatePayment is a helper method to define mock.On call
//   - orderID string
//   - amount float64
func (_e *MockPaymentServiceInterface_Expecter) CreatePayment(orderID interface{}, amount interface{}) *MockPaymentServiceInterface_CreatePayment_Call {
	return &MockPaymentServiceInterface_CreatePayment_Call{Call: _e.mock.On("CreatePayment", orderID, amount)}
}

func (_c *MockPaymentServiceInterface_CreatePayment_Call) Run(run func(orderID string, amount float64)) *MockPaymentServiceInterface_CreatePayment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(float64))
	})
	return _c
}

func (_c *MockPaymentServiceInterface_CreatePayment_Call) Return(_a0 *dtos.CheckoutDTO, _a1 error) *MockPaymentServiceInterface_CreatePayment_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPaymentServiceInterface_CreatePayment_Call) RunAndReturn(run func(string, float64) (*dtos.CheckoutDTO, error)) *MockPaymentServiceInterface_CreatePayment_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPaymentServiceInterface creates a new instance of MockPaymentServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPaymentServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPaymentServiceInterface {
	mock := &MockPaymentServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}