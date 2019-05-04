// Code generated by MockGen. DO NOT EDIT.
// Source: modules/employees/query/query.go

// Package mock_query is a generated GoMock package.
package mock_query

import (
	enum "github.com/azbyluthfan/go-hr/modules/employees/enum"
	model "github.com/azbyluthfan/go-hr/modules/employees/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockEmployeeQuery is a mock of EmployeeQuery interface
type MockEmployeeQuery struct {
	ctrl     *gomock.Controller
	recorder *MockEmployeeQueryMockRecorder
}

// MockEmployeeQueryMockRecorder is the mock recorder for MockEmployeeQuery
type MockEmployeeQueryMockRecorder struct {
	mock *MockEmployeeQuery
}

// NewMockEmployeeQuery creates a new mock instance
func NewMockEmployeeQuery(ctrl *gomock.Controller) *MockEmployeeQuery {
	mock := &MockEmployeeQuery{ctrl: ctrl}
	mock.recorder = &MockEmployeeQueryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEmployeeQuery) EXPECT() *MockEmployeeQueryMockRecorder {
	return m.recorder
}

// GetEmployeeByEmployeeNo mocks base method
func (m *MockEmployeeQuery) GetEmployeeByEmployeeNo(companyId, employeeNo string) (*model.Employee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEmployeeByEmployeeNo", companyId, employeeNo)
	ret0, _ := ret[0].(*model.Employee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEmployeeByEmployeeNo indicates an expected call of GetEmployeeByEmployeeNo
func (mr *MockEmployeeQueryMockRecorder) GetEmployeeByEmployeeNo(companyId, employeeNo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmployeeByEmployeeNo", reflect.TypeOf((*MockEmployeeQuery)(nil).GetEmployeeByEmployeeNo), companyId, employeeNo)
}

// VerifyPassword mocks base method
func (m *MockEmployeeQuery) VerifyPassword(companyId, employeeNo, password string) (*model.Employee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyPassword", companyId, employeeNo, password)
	ret0, _ := ret[0].(*model.Employee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyPassword indicates an expected call of VerifyPassword
func (mr *MockEmployeeQueryMockRecorder) VerifyPassword(companyId, employeeNo, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyPassword", reflect.TypeOf((*MockEmployeeQuery)(nil).VerifyPassword), companyId, employeeNo, password)
}

// CreateNotice mocks base method
func (m *MockEmployeeQuery) CreateNotice(companyId, employeeNo string, noticeType enum.NoticeType, visibility enum.NoticeVisibility, periodStart, periodEnd time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNotice", companyId, employeeNo, noticeType, visibility, periodStart, periodEnd)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateNotice indicates an expected call of CreateNotice
func (mr *MockEmployeeQueryMockRecorder) CreateNotice(companyId, employeeNo, noticeType, visibility, periodStart, periodEnd interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNotice", reflect.TypeOf((*MockEmployeeQuery)(nil).CreateNotice), companyId, employeeNo, noticeType, visibility, periodStart, periodEnd)
}

// GetNotice mocks base method
func (m *MockEmployeeQuery) GetNotice(employeeId string) ([]*model.Notice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNotice", employeeId)
	ret0, _ := ret[0].([]*model.Notice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNotice indicates an expected call of GetNotice
func (mr *MockEmployeeQueryMockRecorder) GetNotice(employeeId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNotice", reflect.TypeOf((*MockEmployeeQuery)(nil).GetNotice), employeeId)
}

// GetCompanyNotice mocks base method
func (m *MockEmployeeQuery) GetCompanyNotice(companyId, visibility string) ([]*model.Notice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCompanyNotice", companyId, visibility)
	ret0, _ := ret[0].([]*model.Notice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCompanyNotice indicates an expected call of GetCompanyNotice
func (mr *MockEmployeeQueryMockRecorder) GetCompanyNotice(companyId, visibility interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCompanyNotice", reflect.TypeOf((*MockEmployeeQuery)(nil).GetCompanyNotice), companyId, visibility)
}
