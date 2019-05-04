package usecase

import (
	"context"
	"github.com/azbyluthfan/go-hr/modules/employees/enum"
	"github.com/azbyluthfan/go-hr/modules/employees/model"
	"time"
)

type EmployeeUseCase interface {
	Login(c context.Context, companyId, employeeNo, password, secretKey string) (*model.AuthResponse, error)
	Hello(c context.Context) (string, error)
	CanCreateNotice(c context.Context, companyId, employeeNo string) error
	CreateNotice(
		c context.Context,
		companyId, employeeNo string,
		noticeType enum.NoticeType,
		visibility enum.NoticeVisibility,
		periodStart, periodEnd time.Time) error
	GetCompanyNotice(c context.Context, companyId string) ([]*model.Notice, error)
	GetNotice(c context.Context) ([]*model.Notice, error)
}
