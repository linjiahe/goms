// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fuwensun/goms/eApi/internal/dao (interfaces: Dao)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	model "github.com/fuwensun/goms/eApi/internal/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockDao is a mock of Dao interface
type MockDao struct {
	ctrl     *gomock.Controller
	recorder *MockDaoMockRecorder
}

// MockDaoMockRecorder is the mock recorder for MockDao
type MockDaoMockRecorder struct {
	mock *MockDao
}

// NewMockDao creates a new mock instance
func NewMockDao(ctrl *gomock.Controller) *MockDao {
	mock := &MockDao{ctrl: ctrl}
	mock.recorder = &MockDaoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDao) EXPECT() *MockDaoMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockDao) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockDaoMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDao)(nil).Close))
}

// CreateUser mocks base method
func (m *MockDao) CreateUser(arg0 context.Context, arg1 *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser
func (mr *MockDaoMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockDao)(nil).CreateUser), arg0, arg1)
}

// CreateUserDB mocks base method
func (m *MockDao) CreateUserDB(arg0 context.Context, arg1 *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserDB", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUserDB indicates an expected call of CreateUserDB
func (mr *MockDaoMockRecorder) CreateUserDB(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserDB", reflect.TypeOf((*MockDao)(nil).CreateUserDB), arg0, arg1)
}

// DelUserCC mocks base method
func (m *MockDao) DelUserCC(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DelUserCC", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DelUserCC indicates an expected call of DelUserCC
func (mr *MockDaoMockRecorder) DelUserCC(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelUserCC", reflect.TypeOf((*MockDao)(nil).DelUserCC), arg0, arg1)
}

// DeleteUser mocks base method
func (m *MockDao) DeleteUser(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser
func (mr *MockDaoMockRecorder) DeleteUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockDao)(nil).DeleteUser), arg0, arg1)
}

// DeleteUserDB mocks base method
func (m *MockDao) DeleteUserDB(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserDB", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserDB indicates an expected call of DeleteUserDB
func (mr *MockDaoMockRecorder) DeleteUserDB(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserDB", reflect.TypeOf((*MockDao)(nil).DeleteUserDB), arg0, arg1)
}

// ExistUserCC mocks base method
func (m *MockDao) ExistUserCC(arg0 context.Context, arg1 int64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExistUserCC", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExistUserCC indicates an expected call of ExistUserCC
func (mr *MockDaoMockRecorder) ExistUserCC(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExistUserCC", reflect.TypeOf((*MockDao)(nil).ExistUserCC), arg0, arg1)
}

// GetUserCC mocks base method
func (m *MockDao) GetUserCC(arg0 context.Context, arg1 int64) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserCC", arg0, arg1)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserCC indicates an expected call of GetUserCC
func (mr *MockDaoMockRecorder) GetUserCC(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserCC", reflect.TypeOf((*MockDao)(nil).GetUserCC), arg0, arg1)
}

// Ping mocks base method
func (m *MockDao) Ping(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping
func (mr *MockDaoMockRecorder) Ping(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockDao)(nil).Ping), arg0)
}

// ReadPingCount mocks base method
func (m *MockDao) ReadPingCount(arg0 context.Context, arg1 model.PingType) (model.PingCount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadPingCount", arg0, arg1)
	ret0, _ := ret[0].(model.PingCount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadPingCount indicates an expected call of ReadPingCount
func (mr *MockDaoMockRecorder) ReadPingCount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadPingCount", reflect.TypeOf((*MockDao)(nil).ReadPingCount), arg0, arg1)
}

// ReadUser mocks base method
func (m *MockDao) ReadUser(arg0 context.Context, arg1 int64) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadUser", arg0, arg1)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadUser indicates an expected call of ReadUser
func (mr *MockDaoMockRecorder) ReadUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadUser", reflect.TypeOf((*MockDao)(nil).ReadUser), arg0, arg1)
}

// ReadUserDB mocks base method
func (m *MockDao) ReadUserDB(arg0 context.Context, arg1 int64) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadUserDB", arg0, arg1)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadUserDB indicates an expected call of ReadUserDB
func (mr *MockDaoMockRecorder) ReadUserDB(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadUserDB", reflect.TypeOf((*MockDao)(nil).ReadUserDB), arg0, arg1)
}

// SetUserCC mocks base method
func (m *MockDao) SetUserCC(arg0 context.Context, arg1 *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUserCC", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetUserCC indicates an expected call of SetUserCC
func (mr *MockDaoMockRecorder) SetUserCC(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUserCC", reflect.TypeOf((*MockDao)(nil).SetUserCC), arg0, arg1)
}

// UpdatePingCount mocks base method
func (m *MockDao) UpdatePingCount(arg0 context.Context, arg1 model.PingType, arg2 model.PingCount) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePingCount", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePingCount indicates an expected call of UpdatePingCount
func (mr *MockDaoMockRecorder) UpdatePingCount(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePingCount", reflect.TypeOf((*MockDao)(nil).UpdatePingCount), arg0, arg1, arg2)
}

// UpdateUser mocks base method
func (m *MockDao) UpdateUser(arg0 context.Context, arg1 *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser
func (mr *MockDaoMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockDao)(nil).UpdateUser), arg0, arg1)
}

// UpdateUserDB mocks base method
func (m *MockDao) UpdateUserDB(arg0 context.Context, arg1 *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserDB", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserDB indicates an expected call of UpdateUserDB
func (mr *MockDaoMockRecorder) UpdateUserDB(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserDB", reflect.TypeOf((*MockDao)(nil).UpdateUserDB), arg0, arg1)
}
