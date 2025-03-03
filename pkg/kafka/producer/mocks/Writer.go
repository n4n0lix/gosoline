// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	context "context"

	kafka "github.com/segmentio/kafka-go"
	mock "github.com/stretchr/testify/mock"
)

// Writer is an autogenerated mock type for the Writer type
type Writer struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *Writer) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stats provides a mock function with given fields:
func (_m *Writer) Stats() kafka.WriterStats {
	ret := _m.Called()

	var r0 kafka.WriterStats
	if rf, ok := ret.Get(0).(func() kafka.WriterStats); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(kafka.WriterStats)
	}

	return r0
}

// WriteMessages provides a mock function with given fields: ctx, msgs
func (_m *Writer) WriteMessages(ctx context.Context, msgs ...kafka.Message) error {
	ret := _m.Called(ctx, msgs)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ...kafka.Message) error); ok {
		r0 = rf(ctx, msgs...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
