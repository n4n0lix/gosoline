// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	types "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	mock "github.com/stretchr/testify/mock"
)

// TransactGetItemBuilder is an autogenerated mock type for the TransactGetItemBuilder type
type TransactGetItemBuilder struct {
	mock.Mock
}

// Build provides a mock function with given fields:
func (_m *TransactGetItemBuilder) Build() (types.TransactGetItem, error) {
	ret := _m.Called()

	var r0 types.TransactGetItem
	if rf, ok := ret.Get(0).(func() types.TransactGetItem); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(types.TransactGetItem)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetItem provides a mock function with given fields:
func (_m *TransactGetItemBuilder) GetItem() interface{} {
	ret := _m.Called()

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

type mockConstructorTestingTNewTransactGetItemBuilder interface {
	mock.TestingT
	Cleanup(func())
}

// NewTransactGetItemBuilder creates a new instance of TransactGetItemBuilder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTransactGetItemBuilder(t mockConstructorTestingTNewTransactGetItemBuilder) *TransactGetItemBuilder {
	mock := &TransactGetItemBuilder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
