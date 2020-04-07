// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/goofinator/usersHttp/internal/repositories (interfaces: Storager)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/goofinator/usersHttp/internal/web/model"
)

// MockStorager is a mock of Storager interface
type MockStorager struct {
	ctrl     *gomock.Controller
	recorder *MockStoragerMockRecorder
}

// MockStoragerMockRecorder is the mock recorder for MockStorager
type MockStoragerMockRecorder struct {
	mock *MockStorager
}

// NewMockStorager creates a new mock instance
func NewMockStorager(ctrl *gomock.Controller) *MockStorager {
	mock := &MockStorager{ctrl: ctrl}
	mock.recorder = &MockStoragerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStorager) EXPECT() *MockStoragerMockRecorder {
	return m.recorder
}

// AddUser mocks base method
func (m *MockStorager) AddUser(arg0 *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUser indicates an expected call of AddUser
func (mr *MockStoragerMockRecorder) AddUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockStorager)(nil).AddUser), arg0)
}

// Close mocks base method
func (m *MockStorager) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockStoragerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStorager)(nil).Close))
}

// DeleteUser mocks base method
func (m *MockStorager) DeleteUser(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser
func (mr *MockStoragerMockRecorder) DeleteUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockStorager)(nil).DeleteUser), arg0)
}

// EditUser mocks base method
func (m *MockStorager) EditUser(arg0 int, arg1 *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EditUser indicates an expected call of EditUser
func (mr *MockStoragerMockRecorder) EditUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditUser", reflect.TypeOf((*MockStorager)(nil).EditUser), arg0, arg1)
}

// GetUsers mocks base method
func (m *MockStorager) GetUsers() ([]*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers")
	ret0, _ := ret[0].([]*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers
func (mr *MockStoragerMockRecorder) GetUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockStorager)(nil).GetUsers))
}