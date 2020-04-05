// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fuwensun/goms/eTest/internal/service (interfaces: Svc)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	model "github.com/fuwensun/goms/eTest/internal/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockSvc is a mock of Svc interface
type MockSvc struct {
	ctrl     *gomock.Controller
	recorder *MockSvcMockRecorder
}

// MockSvcMockRecorder is the mock recorder for MockSvc
type MockSvcMockRecorder struct {
	mock *MockSvc
}

// NewMockSvc creates a new mock instance
func NewMockSvc(ctrl *gomock.Controller) *MockSvc {
	mock := &MockSvc{ctrl: ctrl}
	mock.recorder = &MockSvcMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSvc) EXPECT() *MockSvcMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockSvc) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockSvcMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSvc)(nil).Close))
}

// CreateUser mocks base method
func (m *MockSvc) CreateUser(arg0 context.Context, arg1 *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser
func (mr *MockSvcMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockSvc)(nil).CreateUser), arg0, arg1)
}

// DeleteUser mocks base method
func (m *MockSvc) DeleteUser(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser
func (mr *MockSvcMockRecorder) DeleteUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockSvc)(nil).DeleteUser), arg0, arg1)
}

// Ping mocks base method
func (m *MockSvc) Ping(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping
func (mr *MockSvcMockRecorder) Ping(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockSvc)(nil).Ping), arg0)
}

// ReadGrpcPingCount mocks base method
func (m *MockSvc) ReadGrpcPingCount(arg0 context.Context) (model.PingCount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadGrpcPingCount", arg0)
	ret0, _ := ret[0].(model.PingCount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadGrpcPingCount indicates an expected call of ReadGrpcPingCount
func (mr *MockSvcMockRecorder) ReadGrpcPingCount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadGrpcPingCount", reflect.TypeOf((*MockSvc)(nil).ReadGrpcPingCount), arg0)
}

// ReadHttpPingCount mocks base method
func (m *MockSvc) ReadHttpPingCount(arg0 context.Context) (model.PingCount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadHttpPingCount", arg0)
	ret0, _ := ret[0].(model.PingCount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadHttpPingCount indicates an expected call of ReadHttpPingCount
func (mr *MockSvcMockRecorder) ReadHttpPingCount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadHttpPingCount", reflect.TypeOf((*MockSvc)(nil).ReadHttpPingCount), arg0)
}

// ReadUser mocks base method
func (m *MockSvc) ReadUser(arg0 context.Context, arg1 int64) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadUser", arg0, arg1)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadUser indicates an expected call of ReadUser
func (mr *MockSvcMockRecorder) ReadUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadUser", reflect.TypeOf((*MockSvc)(nil).ReadUser), arg0, arg1)
}

// UpdateGrpcPingCount mocks base method
func (m *MockSvc) UpdateGrpcPingCount(arg0 context.Context, arg1 model.PingCount) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGrpcPingCount", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateGrpcPingCount indicates an expected call of UpdateGrpcPingCount
func (mr *MockSvcMockRecorder) UpdateGrpcPingCount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGrpcPingCount", reflect.TypeOf((*MockSvc)(nil).UpdateGrpcPingCount), arg0, arg1)
}

// UpdateHttpPingCount mocks base method
func (m *MockSvc) UpdateHttpPingCount(arg0 context.Context, arg1 model.PingCount) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateHttpPingCount", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateHttpPingCount indicates an expected call of UpdateHttpPingCount
func (mr *MockSvcMockRecorder) UpdateHttpPingCount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHttpPingCount", reflect.TypeOf((*MockSvc)(nil).UpdateHttpPingCount), arg0, arg1)
}

// UpdateUser mocks base method
func (m *MockSvc) UpdateUser(arg0 context.Context, arg1 *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser
func (mr *MockSvcMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockSvc)(nil).UpdateUser), arg0, arg1)
}