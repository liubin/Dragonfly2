// Code generated by MockGen. DO NOT EDIT.
// Source: d7y.io/dragonfly/v2/cdn/storedriver (interfaces: Driver)

// Package storedriver is a generated GoMock package.
package storedriver

import (
	io "io"
	reflect "reflect"

	unit "d7y.io/dragonfly/v2/pkg/unit"
	gomock "github.com/golang/mock/gomock"
)

// MockDriver is a mock of Driver interface.
type MockDriver struct {
	ctrl     *gomock.Controller
	recorder *MockDriverMockRecorder
}

// MockDriverMockRecorder is the mock recorder for MockDriver.
type MockDriverMockRecorder struct {
	mock *MockDriver
}

// NewMockDriver creates a new mock instance.
func NewMockDriver(ctrl *gomock.Controller) *MockDriver {
	mock := &MockDriver{ctrl: ctrl}
	mock.recorder = &MockDriverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDriver) EXPECT() *MockDriverMockRecorder {
	return m.recorder
}

// CreateBaseDir mocks base method.
func (m *MockDriver) CreateBaseDir() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBaseDir")
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateBaseDir indicates an expected call of CreateBaseDir.
func (mr *MockDriverMockRecorder) CreateBaseDir() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBaseDir", reflect.TypeOf((*MockDriver)(nil).CreateBaseDir))
}

// Exits mocks base method.
func (m *MockDriver) Exits(arg0 *Raw) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exits", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Exits indicates an expected call of Exits.
func (mr *MockDriverMockRecorder) Exits(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exits", reflect.TypeOf((*MockDriver)(nil).Exits), arg0)
}

// Get mocks base method.
func (m *MockDriver) Get(arg0 *Raw) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockDriverMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDriver)(nil).Get), arg0)
}

// GetBaseDir mocks base method.
func (m *MockDriver) GetBaseDir() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBaseDir")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetBaseDir indicates an expected call of GetBaseDir.
func (mr *MockDriverMockRecorder) GetBaseDir() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBaseDir", reflect.TypeOf((*MockDriver)(nil).GetBaseDir))
}

// GetBytes mocks base method.
func (m *MockDriver) GetBytes(arg0 *Raw) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBytes", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBytes indicates an expected call of GetBytes.
func (mr *MockDriverMockRecorder) GetBytes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBytes", reflect.TypeOf((*MockDriver)(nil).GetBytes), arg0)
}

// GetFreeSpace mocks base method.
func (m *MockDriver) GetFreeSpace() (unit.Bytes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFreeSpace")
	ret0, _ := ret[0].(unit.Bytes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFreeSpace indicates an expected call of GetFreeSpace.
func (mr *MockDriverMockRecorder) GetFreeSpace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFreeSpace", reflect.TypeOf((*MockDriver)(nil).GetFreeSpace))
}

// GetPath mocks base method.
func (m *MockDriver) GetPath(arg0 *Raw) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPath", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetPath indicates an expected call of GetPath.
func (mr *MockDriverMockRecorder) GetPath(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPath", reflect.TypeOf((*MockDriver)(nil).GetPath), arg0)
}

// GetTotalAndFreeSpace mocks base method.
func (m *MockDriver) GetTotalAndFreeSpace() (unit.Bytes, unit.Bytes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotalAndFreeSpace")
	ret0, _ := ret[0].(unit.Bytes)
	ret1, _ := ret[1].(unit.Bytes)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTotalAndFreeSpace indicates an expected call of GetTotalAndFreeSpace.
func (mr *MockDriverMockRecorder) GetTotalAndFreeSpace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotalAndFreeSpace", reflect.TypeOf((*MockDriver)(nil).GetTotalAndFreeSpace))
}

// GetTotalSpace mocks base method.
func (m *MockDriver) GetTotalSpace() (unit.Bytes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotalSpace")
	ret0, _ := ret[0].(unit.Bytes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTotalSpace indicates an expected call of GetTotalSpace.
func (mr *MockDriverMockRecorder) GetTotalSpace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotalSpace", reflect.TypeOf((*MockDriver)(nil).GetTotalSpace))
}

// MoveFile mocks base method.
func (m *MockDriver) MoveFile(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MoveFile", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// MoveFile indicates an expected call of MoveFile.
func (mr *MockDriverMockRecorder) MoveFile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MoveFile", reflect.TypeOf((*MockDriver)(nil).MoveFile), arg0, arg1)
}

// Put mocks base method.
func (m *MockDriver) Put(arg0 *Raw, arg1 io.Reader) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Put indicates an expected call of Put.
func (mr *MockDriverMockRecorder) Put(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockDriver)(nil).Put), arg0, arg1)
}

// PutBytes mocks base method.
func (m *MockDriver) PutBytes(arg0 *Raw, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutBytes", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutBytes indicates an expected call of PutBytes.
func (mr *MockDriverMockRecorder) PutBytes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutBytes", reflect.TypeOf((*MockDriver)(nil).PutBytes), arg0, arg1)
}

// Remove mocks base method.
func (m *MockDriver) Remove(arg0 *Raw) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockDriverMockRecorder) Remove(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockDriver)(nil).Remove), arg0)
}

// Stat mocks base method.
func (m *MockDriver) Stat(arg0 *Raw) (*StorageInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stat", arg0)
	ret0, _ := ret[0].(*StorageInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stat indicates an expected call of Stat.
func (mr *MockDriverMockRecorder) Stat(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stat", reflect.TypeOf((*MockDriver)(nil).Stat), arg0)
}

// Walk mocks base method.
func (m *MockDriver) Walk(arg0 *Raw) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Walk", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Walk indicates an expected call of Walk.
func (mr *MockDriverMockRecorder) Walk(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Walk", reflect.TypeOf((*MockDriver)(nil).Walk), arg0)
}
