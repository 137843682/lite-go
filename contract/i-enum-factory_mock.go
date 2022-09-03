// Code generated by MockGen. DO NOT EDIT.
// Source: contract\i-enum-factory.go

// Package contract is a generated GoMock package.
package contract

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIEnumFactory is a mock of IEnumFactory interface.
type MockIEnumFactory struct {
	ctrl     *gomock.Controller
	recorder *MockIEnumFactoryMockRecorder
}

// MockIEnumFactoryMockRecorder is the mock recorder for MockIEnumFactory.
type MockIEnumFactoryMockRecorder struct {
	mock *MockIEnumFactory
}

// NewMockIEnumFactory creates a new mock instance.
func NewMockIEnumFactory(ctrl *gomock.Controller) *MockIEnumFactory {
	mock := &MockIEnumFactory{ctrl: ctrl}
	mock.recorder = &MockIEnumFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIEnumFactory) EXPECT() *MockIEnumFactoryMockRecorder {
	return m.recorder
}

// Build mocks base method.
func (m *MockIEnumFactory) Build(arg0 string) IEnum {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Build", arg0)
	ret0, _ := ret[0].(IEnum)
	return ret0
}

// Build indicates an expected call of Build.
func (mr *MockIEnumFactoryMockRecorder) Build(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockIEnumFactory)(nil).Build), arg0)
}
