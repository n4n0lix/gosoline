// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	kernel "github.com/justtrackio/gosoline/pkg/kernel"
	mock "github.com/stretchr/testify/mock"
)

// Kernel is an autogenerated mock type for the Kernel type
type Kernel struct {
	mock.Mock
}

// Add provides a mock function with given fields: name, moduleFactory, opts
func (_m *Kernel) Add(name string, moduleFactory kernel.ModuleFactory, opts ...kernel.ModuleOption) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, name, moduleFactory)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// AddFactory provides a mock function with given fields: factory
func (_m *Kernel) AddFactory(factory kernel.ModuleMultiFactory) {
	_m.Called(factory)
}

// Run provides a mock function with given fields:
func (_m *Kernel) Run() {
	_m.Called()
}

// Running provides a mock function with given fields:
func (_m *Kernel) Running() <-chan struct{} {
	ret := _m.Called()

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func() <-chan struct{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	return r0
}

// Stop provides a mock function with given fields: reason
func (_m *Kernel) Stop(reason string) {
	_m.Called(reason)
}
