// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fuwensun/goms/eGrpc/api (interfaces: CallClient)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	api "github.com/fuwensun/goms/eGrpc/api"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockCallClient is a mock of CallClient interface
type MockCallClient struct {
	ctrl     *gomock.Controller
	recorder *MockCallClientMockRecorder
}

// MockCallClientMockRecorder is the mock recorder for MockCallClient
type MockCallClientMockRecorder struct {
	mock *MockCallClient
}

// NewMockCallClient creates a new mock instance
func NewMockCallClient(ctrl *gomock.Controller) *MockCallClient {
	mock := &MockCallClient{ctrl: ctrl}
	mock.recorder = &MockCallClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCallClient) EXPECT() *MockCallClientMockRecorder {
	return m.recorder
}

// Ping mocks base method
func (m *MockCallClient) Ping(arg0 context.Context, arg1 *api.Request, arg2 ...grpc.CallOption) (*api.Reply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Ping", varargs...)
	ret0, _ := ret[0].(*api.Reply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Ping indicates an expected call of Ping
func (mr *MockCallClientMockRecorder) Ping(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockCallClient)(nil).Ping), varargs...)
}