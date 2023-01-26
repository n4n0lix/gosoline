// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// Writer is an autogenerated mock type for the Writer type
type Writer struct {
	mock.Mock
}

// Write provides a mock function with given fields: ctx, datetime, items
func (_m *Writer) Write(ctx context.Context, datetime time.Time, items interface{}) error {
	ret := _m.Called(ctx, datetime, items)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, time.Time, interface{}) error); ok {
		r0 = rf(ctx, datetime, items)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WriteToKey provides a mock function with given fields: ctx, key, items
func (_m *Writer) WriteToKey(ctx context.Context, key string, items interface{}) error {
	ret := _m.Called(ctx, key, items)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}) error); ok {
		r0 = rf(ctx, key, items)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewWriter interface {
	mock.TestingT
	Cleanup(func())
}

// NewWriter creates a new instance of Writer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewWriter(t mockConstructorTestingTNewWriter) *Writer {
	mock := &Writer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
