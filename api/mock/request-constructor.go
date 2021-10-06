// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mymmrac/telego/api (interfaces: RequestConstructor)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "github.com/mymmrac/telego/api"
)

// MockRequestConstructor is a mock of RequestConstructor interface.
type MockRequestConstructor struct {
	ctrl     *gomock.Controller
	recorder *MockRequestConstructorMockRecorder
}

// MockRequestConstructorMockRecorder is the mock recorder for MockRequestConstructor.
type MockRequestConstructorMockRecorder struct {
	mock *MockRequestConstructor
}

// NewMockRequestConstructor creates a new mock instance.
func NewMockRequestConstructor(ctrl *gomock.Controller) *MockRequestConstructor {
	mock := &MockRequestConstructor{ctrl: ctrl}
	mock.recorder = &MockRequestConstructorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRequestConstructor) EXPECT() *MockRequestConstructorMockRecorder {
	return m.recorder
}

// JSONRequest mocks base method.
func (m *MockRequestConstructor) JSONRequest(arg0 interface{}) (*api.RequestData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "JSONRequest", arg0)
	ret0, _ := ret[0].(*api.RequestData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// JSONRequest indicates an expected call of JSONRequest.
func (mr *MockRequestConstructorMockRecorder) JSONRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JSONRequest", reflect.TypeOf((*MockRequestConstructor)(nil).JSONRequest), arg0)
}

// MultipartRequest mocks base method.
func (m *MockRequestConstructor) MultipartRequest(arg0 map[string]string, arg1 map[string]api.NamedReader) (*api.RequestData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MultipartRequest", arg0, arg1)
	ret0, _ := ret[0].(*api.RequestData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultipartRequest indicates an expected call of MultipartRequest.
func (mr *MockRequestConstructorMockRecorder) MultipartRequest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultipartRequest", reflect.TypeOf((*MockRequestConstructor)(nil).MultipartRequest), arg0, arg1)
}
