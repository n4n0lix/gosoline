// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	ipread "github.com/justtrackio/gosoline/pkg/ipread"
	mock "github.com/stretchr/testify/mock"
)

// Reader is an autogenerated mock type for the Reader type
type Reader struct {
	mock.Mock
}

// City provides a mock function with given fields: ipString
func (_m *Reader) City(ipString string) (*ipread.GeoCity, error) {
	ret := _m.Called(ipString)

	var r0 *ipread.GeoCity
	if rf, ok := ret.Get(0).(func(string) *ipread.GeoCity); ok {
		r0 = rf(ipString)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ipread.GeoCity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ipString)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewReader interface {
	mock.TestingT
	Cleanup(func())
}

// NewReader creates a new instance of Reader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewReader(t mockConstructorTestingTNewReader) *Reader {
	mock := &Reader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
