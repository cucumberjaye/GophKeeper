// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cucumberjaye/GophKeeper/internal/app/handler/serverhandler (interfaces: KeeperService)

// Package mock_serverhandler is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	models "github.com/cucumberjaye/GophKeeper/internal/app/models"
	gomock "github.com/golang/mock/gomock"
)

// MockKeeperService is a mock of KeeperService interface.
type MockKeeperService struct {
	ctrl     *gomock.Controller
	recorder *MockKeeperServiceMockRecorder
}

// MockKeeperServiceMockRecorder is the mock recorder for MockKeeperService.
type MockKeeperServiceMockRecorder struct {
	mock *MockKeeperService
}

// NewMockKeeperService creates a new mock instance.
func NewMockKeeperService(ctrl *gomock.Controller) *MockKeeperService {
	mock := &MockKeeperService{ctrl: ctrl}
	mock.recorder = &MockKeeperServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKeeperService) EXPECT() *MockKeeperServiceMockRecorder {
	return m.recorder
}

// AddUser mocks base method.
func (m *MockKeeperService) AddUser(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUser indicates an expected call of AddUser.
func (mr *MockKeeperServiceMockRecorder) AddUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockKeeperService)(nil).AddUser), arg0, arg1)
}

// CreateToken mocks base method.
func (m *MockKeeperService) CreateToken(arg0, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateToken", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateToken indicates an expected call of CreateToken.
func (mr *MockKeeperServiceMockRecorder) CreateToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateToken", reflect.TypeOf((*MockKeeperService)(nil).CreateToken), arg0, arg1)
}

// DeleteData mocks base method.
func (m *MockKeeperService) DeleteData(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteData", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteData indicates an expected call of DeleteData.
func (mr *MockKeeperServiceMockRecorder) DeleteData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteData", reflect.TypeOf((*MockKeeperService)(nil).DeleteData), arg0, arg1)
}

// SetBankCardData mocks base method.
func (m *MockKeeperService) SetBankCardData(arg0 string, arg1 models.BankCardData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetBankCardData", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetBankCardData indicates an expected call of SetBankCardData.
func (mr *MockKeeperServiceMockRecorder) SetBankCardData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBankCardData", reflect.TypeOf((*MockKeeperService)(nil).SetBankCardData), arg0, arg1)
}

// SetBinaryData mocks base method.
func (m *MockKeeperService) SetBinaryData(arg0 string, arg1 models.BinaryData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetBinaryData", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetBinaryData indicates an expected call of SetBinaryData.
func (mr *MockKeeperServiceMockRecorder) SetBinaryData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBinaryData", reflect.TypeOf((*MockKeeperService)(nil).SetBinaryData), arg0, arg1)
}

// SetLoginPasswordData mocks base method.
func (m *MockKeeperService) SetLoginPasswordData(arg0 string, arg1 models.LoginPasswordData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetLoginPasswordData", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetLoginPasswordData indicates an expected call of SetLoginPasswordData.
func (mr *MockKeeperServiceMockRecorder) SetLoginPasswordData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLoginPasswordData", reflect.TypeOf((*MockKeeperService)(nil).SetLoginPasswordData), arg0, arg1)
}

// SetTextData mocks base method.
func (m *MockKeeperService) SetTextData(arg0 string, arg1 models.TextData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetTextData", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetTextData indicates an expected call of SetTextData.
func (mr *MockKeeperServiceMockRecorder) SetTextData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTextData", reflect.TypeOf((*MockKeeperService)(nil).SetTextData), arg0, arg1)
}

// Sync mocks base method.
func (m *MockKeeperService) Sync(arg0 string) ([]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sync", arg0)
	ret0, _ := ret[0].([]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sync indicates an expected call of Sync.
func (mr *MockKeeperServiceMockRecorder) Sync(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sync", reflect.TypeOf((*MockKeeperService)(nil).Sync), arg0)
}

// UpdateBankCardData mocks base method.
func (m *MockKeeperService) UpdateBankCardData(arg0 string, arg1 models.BankCardData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBankCardData", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBankCardData indicates an expected call of UpdateBankCardData.
func (mr *MockKeeperServiceMockRecorder) UpdateBankCardData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBankCardData", reflect.TypeOf((*MockKeeperService)(nil).UpdateBankCardData), arg0, arg1)
}

// UpdateBinaryData mocks base method.
func (m *MockKeeperService) UpdateBinaryData(arg0 string, arg1 models.BinaryData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBinaryData", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBinaryData indicates an expected call of UpdateBinaryData.
func (mr *MockKeeperServiceMockRecorder) UpdateBinaryData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBinaryData", reflect.TypeOf((*MockKeeperService)(nil).UpdateBinaryData), arg0, arg1)
}

// UpdateLoginPasswordData mocks base method.
func (m *MockKeeperService) UpdateLoginPasswordData(arg0 string, arg1 models.LoginPasswordData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLoginPasswordData", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateLoginPasswordData indicates an expected call of UpdateLoginPasswordData.
func (mr *MockKeeperServiceMockRecorder) UpdateLoginPasswordData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLoginPasswordData", reflect.TypeOf((*MockKeeperService)(nil).UpdateLoginPasswordData), arg0, arg1)
}

// UpdateTextData mocks base method.
func (m *MockKeeperService) UpdateTextData(arg0 string, arg1 models.TextData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTextData", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTextData indicates an expected call of UpdateTextData.
func (mr *MockKeeperServiceMockRecorder) UpdateTextData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTextData", reflect.TypeOf((*MockKeeperService)(nil).UpdateTextData), arg0, arg1)
}
