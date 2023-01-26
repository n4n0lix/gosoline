// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	http "github.com/justtrackio/gosoline/pkg/http"
	mock "github.com/stretchr/testify/mock"

	nethttp "net/http"

	time "time"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// AddRetryCondition provides a mock function with given fields: f
func (_m *Client) AddRetryCondition(f http.RetryConditionFunc) {
	_m.Called(f)
}

// Delete provides a mock function with given fields: ctx, request
func (_m *Client) Delete(ctx context.Context, request *http.Request) (*http.Response, error) {
	ret := _m.Called(ctx, request)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(context.Context, *http.Request) *http.Response); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *http.Request) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ctx, request
func (_m *Client) Get(ctx context.Context, request *http.Request) (*http.Response, error) {
	ret := _m.Called(ctx, request)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(context.Context, *http.Request) *http.Response); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *http.Request) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewJsonRequest provides a mock function with given fields:
func (_m *Client) NewJsonRequest() *http.Request {
	ret := _m.Called()

	var r0 *http.Request
	if rf, ok := ret.Get(0).(func() *http.Request); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Request)
		}
	}

	return r0
}

// NewRequest provides a mock function with given fields:
func (_m *Client) NewRequest() *http.Request {
	ret := _m.Called()

	var r0 *http.Request
	if rf, ok := ret.Get(0).(func() *http.Request); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Request)
		}
	}

	return r0
}

// NewXmlRequest provides a mock function with given fields:
func (_m *Client) NewXmlRequest() *http.Request {
	ret := _m.Called()

	var r0 *http.Request
	if rf, ok := ret.Get(0).(func() *http.Request); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Request)
		}
	}

	return r0
}

// Patch provides a mock function with given fields: ctx, request
func (_m *Client) Patch(ctx context.Context, request *http.Request) (*http.Response, error) {
	ret := _m.Called(ctx, request)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(context.Context, *http.Request) *http.Response); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *http.Request) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Post provides a mock function with given fields: ctx, request
func (_m *Client) Post(ctx context.Context, request *http.Request) (*http.Response, error) {
	ret := _m.Called(ctx, request)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(context.Context, *http.Request) *http.Response); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *http.Request) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Put provides a mock function with given fields: ctx, request
func (_m *Client) Put(ctx context.Context, request *http.Request) (*http.Response, error) {
	ret := _m.Called(ctx, request)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(context.Context, *http.Request) *http.Response); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *http.Request) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetCookie provides a mock function with given fields: c
func (_m *Client) SetCookie(c *nethttp.Cookie) {
	_m.Called(c)
}

// SetCookies provides a mock function with given fields: cs
func (_m *Client) SetCookies(cs []*nethttp.Cookie) {
	_m.Called(cs)
}

// SetProxyUrl provides a mock function with given fields: p
func (_m *Client) SetProxyUrl(p string) {
	_m.Called(p)
}

// SetRedirectValidator provides a mock function with given fields: allowRequest
func (_m *Client) SetRedirectValidator(allowRequest func(*nethttp.Request) bool) {
	_m.Called(allowRequest)
}

// SetTimeout provides a mock function with given fields: timeout
func (_m *Client) SetTimeout(timeout time.Duration) {
	_m.Called(timeout)
}

// SetUserAgent provides a mock function with given fields: ua
func (_m *Client) SetUserAgent(ua string) {
	_m.Called(ua)
}

type mockConstructorTestingTNewClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClient(t mockConstructorTestingTNewClient) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
