// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// TypedModule is an autogenerated mock type for the TypedModule type
type TypedModule struct {
	mock.Mock
}

// IsBackground provides a mock function with given fields:
func (_m *TypedModule) IsBackground() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsEssential provides a mock function with given fields:
func (_m *TypedModule) IsEssential() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type mockConstructorTestingTNewTypedModule interface {
	mock.TestingT
	Cleanup(func())
}

// NewTypedModule creates a new instance of TypedModule. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTypedModule(t mockConstructorTestingTNewTypedModule) *TypedModule {
	mock := &TypedModule{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
