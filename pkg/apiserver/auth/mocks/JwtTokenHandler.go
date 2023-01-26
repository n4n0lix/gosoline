// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	jwt "github.com/dgrijalva/jwt-go"
	auth "github.com/justtrackio/gosoline/pkg/apiserver/auth"

	mock "github.com/stretchr/testify/mock"
)

// JwtTokenHandler is an autogenerated mock type for the JwtTokenHandler type
type JwtTokenHandler struct {
	mock.Mock
}

// Sign provides a mock function with given fields: user
func (_m *JwtTokenHandler) Sign(user auth.SignUserInput) (*string, error) {
	ret := _m.Called(user)

	var r0 *string
	if rf, ok := ret.Get(0).(func(auth.SignUserInput) *string); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(auth.SignUserInput) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Valid provides a mock function with given fields: jwtToken
func (_m *JwtTokenHandler) Valid(jwtToken string) (bool, *jwt.Token, error) {
	ret := _m.Called(jwtToken)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(jwtToken)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 *jwt.Token
	if rf, ok := ret.Get(1).(func(string) *jwt.Token); ok {
		r1 = rf(jwtToken)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*jwt.Token)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(jwtToken)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type mockConstructorTestingTNewJwtTokenHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewJwtTokenHandler creates a new instance of JwtTokenHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewJwtTokenHandler(t mockConstructorTestingTNewJwtTokenHandler) *JwtTokenHandler {
	mock := &JwtTokenHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
