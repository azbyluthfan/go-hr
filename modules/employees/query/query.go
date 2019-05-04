package query

import (
	"github.com/azbyluthfan/go-hr/modules/employees/enum"
	"github.com/azbyluthfan/go-hr/modules/employees/model"
	"time"
)

type EmployeeQuery interface {
	GetEmployeeByEmployeeNo(companyId, employeeNo string) (*model.Employee, error)
	VerifyPassword(companyId, employeeNo, password string) (*model.Employee, error)
	CreateNotice(
		companyId, employeeNo string,
		noticeType enum.NoticeType,
		visibility enum.NoticeVisibility,
		periodStart, periodEnd time.Time) error
	GetNotice(employeeId string) ([]*model.Notice, error)
	GetCompanyNotice(companyId, visibility string) ([]*model.Notice, error)
}
