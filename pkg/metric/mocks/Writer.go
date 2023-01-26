// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	metric "github.com/justtrackio/gosoline/pkg/metric"
	mock "github.com/stretchr/testify/mock"
)

// Writer is an autogenerated mock type for the Writer type
type Writer struct {
	mock.Mock
}

// GetPriority provides a mock function with given fields:
func (_m *Writer) GetPriority() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Write provides a mock function with given fields: batch
func (_m *Writer) Write(batch metric.Data) {
	_m.Called(batch)
}

// WriteOne provides a mock function with given fields: data
func (_m *Writer) WriteOne(data *metric.Datum) {
	_m.Called(data)
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
