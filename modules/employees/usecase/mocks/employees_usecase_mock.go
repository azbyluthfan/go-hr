// Code generated by MockGen. DO NOT EDIT.
// Source: modules/employees/usecase/usecase.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	enum "github.com/azbyluthfan/go-hr/modules/employees/enum"
	model "github.com/azbyluthfan/go-hr/modules/employees/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockEmployeeUseCase is a mock of EmployeeUseCase interface
type MockEmployeeUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockEmployeeUseCaseMockRecorder
}

// MockEmployeeUseCaseMockRecorder is the mock recorder for MockEmployeeUseCase
type MockEmployeeUseCaseMockRecorder struct {
	mock *MockEmployeeUseCase
}

// NewMockEmployeeUseCase creates a new mock instance
func NewMockEmployeeUseCase(ctrl *gomock.Controller) *MockEmployeeUseCase {
	mock := &MockEmployeeUseCase{ctrl: ctrl}
	mock.recorder = &MockEmployeeUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEmployeeUseCase) EXPECT() *MockEmployeeUseCaseMockRecorder {
	return m.recorder
}

// Login mocks base method
func (m *MockEmployeeUseCase) Login(c context.Context, companyId, employeeNo, password, secretKey string) (*model.AuthResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", c, companyId, employeeNo, password, secretKey)
	ret0, _ := ret[0].(*model.AuthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login
func (mr *MockEmployeeUseCaseMockRecorder) Login(c, companyId, employeeNo, password, secretKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockEmployeeUseCase)(nil).Login), c, companyId, employeeNo, password, secretKey)
}

// Hello mocks base method
func (m *MockEmployeeUseCase) Hello(c context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hello", c)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Hello indicates an expected call of Hello
func (mr *MockEmployeeUseCaseMockRecorder) Hello(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hello", reflect.TypeOf((*MockEmployeeUseCase)(nil).Hello), c)
}

// CanCreateNotice mocks base method
func (m *MockEmployeeUseCase) CanCreateNotice(c context.Context, companyId, employeeNo string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanCreateNotice", c, companyId, employeeNo)
	ret0, _ := ret[0].(error)
	return ret0
}

// CanCreateNotice indicates an expected call of CanCreateNotice
func (mr *MockEmployeeUseCaseMockRecorder) CanCreateNotice(c, companyId, employeeNo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanCreateNotice", reflect.TypeOf((*MockEmployeeUseCase)(nil).CanCreateNotice), c, companyId, employeeNo)
}

// CreateNotice mocks base method
func (m *MockEmployeeUseCase) CreateNotice(c context.Context, companyId, employeeNo string, noticeType enum.NoticeType, visibility enum.NoticeVisibility, periodStart, periodEnd time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNotice", c, companyId, employeeNo, noticeType, visibility, periodStart, periodEnd)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateNotice indicates an expected call of CreateNotice
func (mr *MockEmployeeUseCaseMockRecorder) CreateNotice(c, companyId, employeeNo, noticeType, visibility, periodStart, periodEnd interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNotice", reflect.TypeOf((*MockEmployeeUseCase)(nil).CreateNotice), c, companyId, employeeNo, noticeType, visibility, periodStart, periodEnd)
}

// GetCompanyNotice mocks base method
func (m *MockEmployeeUseCase) GetCompanyNotice(c context.Context, companyId string) ([]*model.Notice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCompanyNotice", c, companyId)
	ret0, _ := ret[0].([]*model.Notice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCompanyNotice indicates an expected call of GetCompanyNotice
func (mr *MockEmployeeUseCaseMockRecorder) GetCompanyNotice(c, companyId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCompanyNotice", reflect.TypeOf((*MockEmployeeUseCase)(nil).GetCompanyNotice), c, companyId)
}

// GetNotice mocks base method
func (m *MockEmployeeUseCase) GetNotice(c context.Context) ([]*model.Notice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNotice", c)
	ret0, _ := ret[0].([]*model.Notice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNotice indicates an expected call of GetNotice
func (mr *MockEmployeeUseCaseMockRecorder) GetNotice(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNotice", reflect.TypeOf((*MockEmployeeUseCase)(nil).GetNotice), c)
}