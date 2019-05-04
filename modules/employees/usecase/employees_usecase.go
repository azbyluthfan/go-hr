package usecase

import (
	"context"
	"errors"
	token "github.com/azbyluthfan/go-hr/helper/token"
	"github.com/azbyluthfan/go-hr/modules/employees/enum"
	"github.com/azbyluthfan/go-hr/modules/employees/model"
	"github.com/azbyluthfan/go-hr/modules/employees/query"
	"github.com/dgrijalva/jwt-go"
	"github.com/rs/xid"
	"time"
)

type employeeUseCaseImpl struct {
	employeeQuery query.EmployeeQuery
}

// Create a EmployeeUseCase instance
func NewEmployeeUseCase(eq query.EmployeeQuery) EmployeeUseCase {
	return &employeeUseCaseImpl{
		employeeQuery: eq,
	}
}

// Login for specified employee.
// Create an AccessToken if password is verified.
func (impl *employeeUseCaseImpl) Login(c context.Context, companyId, employeeNo, password, secretKey string) (*model.AuthResponse, error) {

	employee, err := impl.employeeQuery.VerifyPassword(companyId, employeeNo, password)
	if err != nil {
		return nil, err
	}

	// Create the token
	guid := xid.New()
	duration, _ := time.ParseDuration("1h")

	claim := jwt.StandardClaims{
		Audience: "www.go.hr",
		ExpiresAt: time.Now().Add(duration).Unix(),
		Id: guid.String(),
	}

	tokenGen := token.NewJwtGenerator(secretKey, duration, duration)
	accessToken := <-tokenGen.GenerateAccessToken(token.BearerClaims{
		employee.EmployeeNo,
		employee.ID,
		employee.Role.String(),
		employee.CompanyID,
		true,
		claim,
	})

	if accessToken.Error != nil {
		return nil, accessToken.Error
	}

	return &model.AuthResponse{
		Token: accessToken.AccessToken,
	}, nil
}

// Return hello message for logged in employee
func (impl *employeeUseCaseImpl) Hello(c context.Context) (string, error) {
	employeeNo := c.Value("employeeNo")
	if employeeNo == nil {
		return "", errors.New("Failed in parsing employeeNo")
	}

	return "Hello, " + employeeNo.(string) + "!", nil
}

// Check if logged in user can create notice for selected employee
func (impl *employeeUseCaseImpl) CanCreateNotice(
	c context.Context,
	companyId, employeeNo string) error {

	if c.Value("companyId") == nil || c.Value("employeeNo") == nil || c.Value("role") == nil {
		return errors.New("Failed in parsing employeeNo, companyId, or role")
	}

	if c.Value("companyId").(string) != companyId &&
		c.Value("employeeNo").(string) != employeeNo &&
		c.Value("role").(string) != enum.ADMIN.String() {
		return errors.New("Have no privilege to create notice for other employee")
	}

	return nil
}


// Create notice for a Employee
func (impl *employeeUseCaseImpl) CreateNotice(
	c context.Context,
	companyId, employeeNo string,
	noticeType enum.NoticeType,
	visibility enum.NoticeVisibility,
	periodStart, periodEnd time.Time) error {

	return impl.employeeQuery.CreateNotice(companyId, employeeNo, noticeType, visibility, periodStart, periodEnd)
}

// Get list of Notice from employee of selected company
func (impl *employeeUseCaseImpl) GetCompanyNotice(c context.Context, companyId string) ([]*model.Notice, error) {

	if c.Value("companyId") == nil || c.Value("role") == nil {
		return nil, errors.New("Failed in parsing employeeNo or role")
	}
	if c.Value("role").(string) == enum.ADMIN.String() {
		return impl.employeeQuery.GetCompanyNotice(c.Value("companyId").(string), "")
	}
	return impl.employeeQuery.GetCompanyNotice(c.Value("companyId").(string), enum.PUBLIC.String())
}

// Get list of Notice from logged in user
func (impl *employeeUseCaseImpl) GetNotice(c context.Context) ([]*model.Notice, error) {

	if c.Value("employeeId") == nil {
		return nil, errors.New("Failed in parsing employeeId")
	}
	return impl.employeeQuery.GetNotice(c.Value("employeeId").(string))
}
