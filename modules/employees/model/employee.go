package model

import (
	"github.com/go-sql-driver/mysql"
	//"gopkg.in/go-playground/validator.v8"
	"github.com/azbyluthfan/go-hr/helper/token"
	"github.com/azbyluthfan/go-hr/modules/employees/enum"
)


type Company struct {
	ID		string	`db:"id"`
	Name	string	`db:"name"`
}

type Employee struct {
	ID				string				`db:"id"`
	Name 			string 				`db:"name"`
	CompanyID		string				`db:"company_id"`
	Role			enum.EmployeeRole	`db:"role"`
	EmployeeNo		string				`db:"employee_no"`
	Password		string				`db:"password"`
	FailedLogin 	int					`db:"failed_login_count"`
	FailedLoginTime	mysql.NullTime		`db:"failed_login_time"`
	JailTime		mysql.NullTime		`db:"jail_time"`
}

type Notice struct {
	ID			string					`db:"id"`
	EmployeeID	string					`db:"employee_id"`
	Type 		enum.NoticeType 		`db:"type"`
	Visibility	enum.NoticeVisibility	`db:"visibility"`
	PeriodStart mysql.NullTime			`db:"period_start" time_format="2000-12-31"`
	PeriodEnd	mysql.NullTime			`db:"period_end" time_format="2000-12-31"`
}

type LoginParam struct {
	CompanyID	string 	`json:"companyID" binding:"required"`
	EmployeeNo	string 	`json:"employeeNo" binding:"required"`
	Password	string  `json:"password" binding:"required"`
}

type AuthResponse struct {
	Token token.AccessToken `json:"token"`
}

type CreateNoticeParam struct {
	CompanyID	string					`json:"companyID" binding:"required"`
	EmployeeNo	string					`json:"employeeNo" binding:"required"`
	Type		enum.NoticeType			`json:"type" binding:"exists"`
	Visibility	enum.NoticeVisibility	`json:"visibility" binding:"exists"`
	PeriodStart string					`json:"periodStart" binding:"required"`
	PeriodEnd	string					`json:"periodEnd" binding:"required"`
}
