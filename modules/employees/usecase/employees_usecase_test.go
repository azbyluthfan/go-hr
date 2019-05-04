package usecase

import (
	"context"
	"github.com/azbyluthfan/go-hr/modules/employees/enum"
	"github.com/azbyluthfan/go-hr/modules/employees/model"
	mock_query "github.com/azbyluthfan/go-hr/modules/employees/query/mocks"
	"github.com/dgrijalva/jwt-go"
	"github.com/golang/mock/gomock"
	"strings"
	"testing"
	"time"
)

var secret = "this-is-secret"

func TestEmployeeUseCaseImpl_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	eQuery := mock_query.NewMockEmployeeQuery(ctrl)

	eUseCase := NewEmployeeUseCase(eQuery)
	employee := model.Employee{ID: "1", Role: enum.ADMIN, CompanyID: "2", EmployeeNo: "x", Password: "123456"}

	eQuery.EXPECT().VerifyPassword(gomock.Any(), gomock.Any(), gomock.Any()).Return(&employee, nil)
	res, _ := eUseCase.Login(context.TODO(), employee.CompanyID, employee.EmployeeNo, employee.Password, secret)

	token, _ := jwt.Parse(res.Token.AccessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if token.Claims.(jwt.MapClaims)["eno"].(string) != employee.EmployeeNo {
		t.Errorf("Incorrect employee no claim, got: %s, wants: %s", token.Claims.(jwt.MapClaims)["eno"].(string), employee.EmployeeNo)
	}
}

func TestEmployeeUseCaseImpl_Hello(t *testing.T) {
	ctrl := gomock.NewController(t)
	eQuery := mock_query.NewMockEmployeeQuery(ctrl)

	eUseCase := NewEmployeeUseCase(eQuery)

	t.Run("Context is populated with employeeNo key", func(t *testing.T) {

		c := context.WithValue(context.Background(), "employeeNo", "x")
		res, _ := eUseCase.Hello(c)

		if res != "Hello, " + c.Value("employeeNo").(string) + "!" {
			t.Errorf("Incorrect return message, got: %s", res)
		}
	})

	t.Run("Should return error", func(t *testing.T) {

		c := context.Background()
		_, err := eUseCase.Hello(c)

		if err == nil {
			t.Errorf("Should return error")
		}

	})
}

func TestEmployeeUseCaseImpl_CanCreateNotice(t *testing.T) {
	ctrl := gomock.NewController(t)
	eQuery := mock_query.NewMockEmployeeQuery(ctrl)

	eUseCase := NewEmployeeUseCase(eQuery)

	t.Run("Should return error", func(t *testing.T) {

		c := context.Background()
		err := eUseCase.CanCreateNotice(c, "", "")

		if err == nil {
			t.Errorf("Should return error")
		}

	})
	t.Run("Should return error no privilege", func(t *testing.T) {
		c := context.WithValue(context.Background(), "employeeNo", "x")
		c = context.WithValue(c, "companyId", "1")
		c = context.WithValue(c, "role", enum.NORMAL.String())

		err := eUseCase.CanCreateNotice(c, "a", "b")

		if err == nil || (err != nil && !strings.Contains(err.Error(), "no privilege")) {
			t.Errorf("Should return error no privilege %v", err)
		}
	})

}

func TestEmployeeUseCaseImpl_CreateNotice(t *testing.T) {
	ctrl := gomock.NewController(t)
	eQuery := mock_query.NewMockEmployeeQuery(ctrl)

	eUseCase := NewEmployeeUseCase(eQuery)

	eQuery.EXPECT().CreateNotice(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	err := eUseCase.CreateNotice(context.TODO(), "", "", enum.REMOTE, enum.PRIVATE, time.Now(), time.Now())

	if err != nil {
		t.Errorf("Should return no error, got %s error", err.Error())
	}
}

func TestEmployeeUseCaseImpl_GetCompanyNotice(t *testing.T) {
	ctrl := gomock.NewController(t)
	eQuery := mock_query.NewMockEmployeeQuery(ctrl)

	eUseCase := NewEmployeeUseCase(eQuery)
	c := context.WithValue(context.Background(), "companyId", "1")

	t.Run("Should return company notice for admin", func(t *testing.T) {
		c = context.WithValue(c, "role", "ADMIN")

		eQuery.EXPECT().GetCompanyNotice(gomock.Any(), gomock.Any()).Return(nil, nil)
		_, err := eUseCase.GetCompanyNotice(c, "1")

		if err != nil {
			t.Errorf("Should return no error, got %s error", err.Error())
		}
	})
	t.Run("Should return company notice for employee", func(t *testing.T) {
		c = context.WithValue(c, "role", "NORMAL")

		eQuery.EXPECT().GetCompanyNotice(gomock.Any(), gomock.Any()).Return(nil, nil)
		_, err := eUseCase.GetCompanyNotice(c, "1")

		if err != nil {
			t.Errorf("Should return no error, got %s error", err.Error())
		}
	})

}

func TestEmployeeUseCaseImpl_GetNotice(t *testing.T) {
	ctrl := gomock.NewController(t)
	eQuery := mock_query.NewMockEmployeeQuery(ctrl)

	eUseCase := NewEmployeeUseCase(eQuery)

	t.Run("Should return error", func(t *testing.T) {

		eQuery.EXPECT().GetNotice(gomock.Any()).Return(nil, nil)
		_, err := eUseCase.GetNotice(context.TODO())

		if err == nil || (err != nil && !strings.Contains(err.Error(), "Failed in parsing")) {
			t.Errorf("Should return error Failed in parsing %v", err)
		}
	})
	t.Run("Should not return error", func(t *testing.T) {
		c := context.WithValue(context.Background(), "employeeId", "1")

		eQuery.EXPECT().GetNotice(gomock.Any()).Return(nil, nil)
		_, err := eUseCase.GetNotice(c)

		if err != nil {
			t.Errorf("Should not return error %v", err.Error())
		}
	})

}