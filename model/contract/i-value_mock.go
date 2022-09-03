// Code generated by MockGen. DO NOT EDIT.
// Source: model\contract\i-value.go

// Package contract is a generated GoMock package.
package contract

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIValue is a mock of IValue interface.
type MockIValue struct {
	ctrl     *gomock.Controller
	recorder *MockIValueMockRecorder
}

// MockIValueMockRecorder is the mock recorder for MockIValue.
type MockIValueMockRecorder struct {
	mock *MockIValue
}

// NewMockIValue creates a new mock instance.
func NewMockIValue(ctrl *gomock.Controller) *MockIValue {
	mock := &MockIValue{ctrl: ctrl}
	mock.recorder = &MockIValueMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIValue) EXPECT() *MockIValueMockRecorder {
	return m.recorder
}

// GetCount mocks base method.
func (m *MockIValue) GetCount() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCount")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetCount indicates an expected call of GetCount.
func (mr *MockIValueMockRecorder) GetCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCount", reflect.TypeOf((*MockIValue)(nil).GetCount))
}

// GetSource mocks base method.
func (m *MockIValue) GetSource() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSource")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetSource indicates an expected call of GetSource.
func (mr *MockIValueMockRecorder) GetSource() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSource", reflect.TypeOf((*MockIValue)(nil).GetSource))
}

// GetValueType mocks base method.
func (m *MockIValue) GetValueType() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValueType")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetValueType indicates an expected call of GetValueType.
func (mr *MockIValueMockRecorder) GetValueType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValueType", reflect.TypeOf((*MockIValue)(nil).GetValueType))
}
