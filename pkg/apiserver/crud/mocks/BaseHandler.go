// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	crud "github.com/justtrackio/gosoline/pkg/apiserver/crud"
	db_repo "github.com/justtrackio/gosoline/pkg/db-repo"

	mock "github.com/stretchr/testify/mock"
)

// BaseHandler is an autogenerated mock type for the BaseHandler type
type BaseHandler struct {
	mock.Mock
}

// GetModel provides a mock function with given fields:
func (_m *BaseHandler) GetModel() db_repo.ModelBased {
	ret := _m.Called()

	var r0 db_repo.ModelBased
	if rf, ok := ret.Get(0).(func() db_repo.ModelBased); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(db_repo.ModelBased)
		}
	}

	return r0
}

// GetRepository provides a mock function with given fields:
func (_m *BaseHandler) GetRepository() crud.Repository {
	ret := _m.Called()

	var r0 crud.Repository
	if rf, ok := ret.Get(0).(func() crud.Repository); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(crud.Repository)
		}
	}

	return r0
}

// TransformOutput provides a mock function with given fields: ctx, model, apiView
func (_m *BaseHandler) TransformOutput(ctx context.Context, model db_repo.ModelBased, apiView string) (interface{}, error) {
	ret := _m.Called(ctx, model, apiView)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(context.Context, db_repo.ModelBased, string) interface{}); ok {
		r0 = rf(ctx, model, apiView)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, db_repo.ModelBased, string) error); ok {
		r1 = rf(ctx, model, apiView)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewBaseHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewBaseHandler creates a new instance of BaseHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBaseHandler(t mockConstructorTestingTNewBaseHandler) *BaseHandler {
	mock := &BaseHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
