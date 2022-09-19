// Code generated by MockGen. DO NOT EDIT.
// Source: contract\i-traceable.go

// Package contract is a generated GoMock package.
package contract

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockITraceable is a mock of ITraceable interface.
type MockITraceable struct {
	ctrl     *gomock.Controller
	recorder *MockITraceableMockRecorder
}

// MockITraceableMockRecorder is the mock recorder for MockITraceable.
type MockITraceableMockRecorder struct {
	mock *MockITraceable
}

// NewMockITraceable creates a new mock instance.
func NewMockITraceable(ctrl *gomock.Controller) *MockITraceable {
	mock := &MockITraceable{ctrl: ctrl}
	mock.recorder = &MockITraceableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITraceable) EXPECT() *MockITraceableMockRecorder {
	return m.recorder
}

// WithContext mocks base method.
func (m *MockITraceable) WithContext(ctx context.Context) reflect.Value {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithContext", ctx)
	ret0, _ := ret[0].(reflect.Value)
	return ret0
}

// WithContext indicates an expected call of WithContext.
func (mr *MockITraceableMockRecorder) WithContext(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithContext", reflect.TypeOf((*MockITraceable)(nil).WithContext), ctx)
}
