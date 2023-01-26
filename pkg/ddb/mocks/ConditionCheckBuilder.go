// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	expression "github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	ddb "github.com/justtrackio/gosoline/pkg/ddb"
	mock "github.com/stretchr/testify/mock"

	types "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// ConditionCheckBuilder is an autogenerated mock type for the ConditionCheckBuilder type
type ConditionCheckBuilder struct {
	mock.Mock
}

// Build provides a mock function with given fields: result
func (_m *ConditionCheckBuilder) Build(result interface{}) (*types.ConditionCheck, error) {
	ret := _m.Called(result)

	var r0 *types.ConditionCheck
	if rf, ok := ret.Get(0).(func(interface{}) *types.ConditionCheck); ok {
		r0 = rf(result)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ConditionCheck)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(result)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReturnAllOld provides a mock function with given fields:
func (_m *ConditionCheckBuilder) ReturnAllOld() ddb.ConditionCheckBuilder {
	ret := _m.Called()

	var r0 ddb.ConditionCheckBuilder
	if rf, ok := ret.Get(0).(func() ddb.ConditionCheckBuilder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.ConditionCheckBuilder)
		}
	}

	return r0
}

// ReturnNone provides a mock function with given fields:
func (_m *ConditionCheckBuilder) ReturnNone() ddb.ConditionCheckBuilder {
	ret := _m.Called()

	var r0 ddb.ConditionCheckBuilder
	if rf, ok := ret.Get(0).(func() ddb.ConditionCheckBuilder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.ConditionCheckBuilder)
		}
	}

	return r0
}

// WithCondition provides a mock function with given fields: cond
func (_m *ConditionCheckBuilder) WithCondition(cond expression.ConditionBuilder) ddb.ConditionCheckBuilder {
	ret := _m.Called(cond)

	var r0 ddb.ConditionCheckBuilder
	if rf, ok := ret.Get(0).(func(expression.ConditionBuilder) ddb.ConditionCheckBuilder); ok {
		r0 = rf(cond)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.ConditionCheckBuilder)
		}
	}

	return r0
}

// WithHash provides a mock function with given fields: hashValue
func (_m *ConditionCheckBuilder) WithHash(hashValue interface{}) ddb.ConditionCheckBuilder {
	ret := _m.Called(hashValue)

	var r0 ddb.ConditionCheckBuilder
	if rf, ok := ret.Get(0).(func(interface{}) ddb.ConditionCheckBuilder); ok {
		r0 = rf(hashValue)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.ConditionCheckBuilder)
		}
	}

	return r0
}

// WithKeys provides a mock function with given fields: keys
func (_m *ConditionCheckBuilder) WithKeys(keys ...interface{}) ddb.ConditionCheckBuilder {
	var _ca []interface{}
	_ca = append(_ca, keys...)
	ret := _m.Called(_ca...)

	var r0 ddb.ConditionCheckBuilder
	if rf, ok := ret.Get(0).(func(...interface{}) ddb.ConditionCheckBuilder); ok {
		r0 = rf(keys...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.ConditionCheckBuilder)
		}
	}

	return r0
}

// WithRange provides a mock function with given fields: rangeValue
func (_m *ConditionCheckBuilder) WithRange(rangeValue interface{}) ddb.ConditionCheckBuilder {
	ret := _m.Called(rangeValue)

	var r0 ddb.ConditionCheckBuilder
	if rf, ok := ret.Get(0).(func(interface{}) ddb.ConditionCheckBuilder); ok {
		r0 = rf(rangeValue)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.ConditionCheckBuilder)
		}
	}

	return r0
}

type mockConstructorTestingTNewConditionCheckBuilder interface {
	mock.TestingT
	Cleanup(func())
}

// NewConditionCheckBuilder creates a new instance of ConditionCheckBuilder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewConditionCheckBuilder(t mockConstructorTestingTNewConditionCheckBuilder) *ConditionCheckBuilder {
	mock := &ConditionCheckBuilder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
