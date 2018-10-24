// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ory/fosite (interfaces: Hasher)

package internal

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHasher is a mock of Hasher interface
type MockHasher struct {
	ctrl     *gomock.Controller
	recorder *MockHasherMockRecorder
}

// MockHasherMockRecorder is the mock recorder for MockHasher
type MockHasherMockRecorder struct {
	mock *MockHasher
}

// NewMockHasher creates a new mock instance
func NewMockHasher(ctrl *gomock.Controller) *MockHasher {
	mock := &MockHasher{ctrl: ctrl}
	mock.recorder = &MockHasherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockHasher) EXPECT() *MockHasherMockRecorder {
	return _m.recorder
}

// Compare mocks base method
func (_m *MockHasher) Compare(_param0 context.Context, _param1 []byte, _param2 []byte) error {
	ret := _m.ctrl.Call(_m, "Compare", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Compare indicates an expected call of Compare
func (_mr *MockHasherMockRecorder) Compare(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Compare", reflect.TypeOf((*MockHasher)(nil).Compare), arg0, arg1, arg2)
}

// Hash mocks base method
func (_m *MockHasher) Hash(_param0 context.Context, _param1 []byte) ([]byte, error) {
	ret := _m.ctrl.Call(_m, "Hash", _param0, _param1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Hash indicates an expected call of Hash
func (_mr *MockHasherMockRecorder) Hash(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Hash", reflect.TypeOf((*MockHasher)(nil).Hash), arg0, arg1)
}
