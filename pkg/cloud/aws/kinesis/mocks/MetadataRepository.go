// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	kinesis "github.com/justtrackio/gosoline/pkg/cloud/aws/kinesis"
	mock "github.com/stretchr/testify/mock"
)

// MetadataRepository is an autogenerated mock type for the MetadataRepository type
type MetadataRepository struct {
	mock.Mock
}

// AcquireShard provides a mock function with given fields: ctx, shardId
func (_m *MetadataRepository) AcquireShard(ctx context.Context, shardId kinesis.ShardId) (kinesis.Checkpoint, error) {
	ret := _m.Called(ctx, shardId)

	var r0 kinesis.Checkpoint
	if rf, ok := ret.Get(0).(func(context.Context, kinesis.ShardId) kinesis.Checkpoint); ok {
		r0 = rf(ctx, shardId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(kinesis.Checkpoint)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, kinesis.ShardId) error); ok {
		r1 = rf(ctx, shardId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeregisterClient provides a mock function with given fields: ctx
func (_m *MetadataRepository) DeregisterClient(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IsShardFinished provides a mock function with given fields: ctx, shardId
func (_m *MetadataRepository) IsShardFinished(ctx context.Context, shardId kinesis.ShardId) (bool, error) {
	ret := _m.Called(ctx, shardId)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, kinesis.ShardId) bool); ok {
		r0 = rf(ctx, shardId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, kinesis.ShardId) error); ok {
		r1 = rf(ctx, shardId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterClient provides a mock function with given fields: ctx
func (_m *MetadataRepository) RegisterClient(ctx context.Context) (int, int, error) {
	ret := _m.Called(ctx)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context) int); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context) int); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context) error); ok {
		r2 = rf(ctx)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type mockConstructorTestingTNewMetadataRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewMetadataRepository creates a new instance of MetadataRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMetadataRepository(t mockConstructorTestingTNewMetadataRepository) *MetadataRepository {
	mock := &MetadataRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
