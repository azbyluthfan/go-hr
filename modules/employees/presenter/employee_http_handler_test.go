package presenter

import (
	"encoding/json"
	"errors"
	"github.com/azbyluthfan/go-hr/helper/token"
	"github.com/azbyluthfan/go-hr/modules/employees/enum"
	"github.com/azbyluthfan/go-hr/modules/employees/model"
	mock_usecase "github.com/azbyluthfan/go-hr/modules/employees/usecase/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestEmployeeHttpHandler_Login(t *testing.T) {

	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	eUseCase := mock_usecase.NewMockEmployeeUseCase(ctrl)
	eHandler := NewEmployeeHttpHandler(eUseCase)

	router := gin.New()
	router.POST("/auth/login", eHandler.Login)

	t.Run("Successful login", func(t *testing.T) {

		loginParam := model.LoginParam{CompanyID: "1", EmployeeNo: "2", Password: "3"}
		param, _ := json.Marshal(loginParam)

		req := httptest.NewRequest("POST", "/auth/login", strings.NewReader(string(param)))
		resp := httptest.NewRecorder()

		authResp := model.AuthResponse{Token: token.AccessToken{}}
		eUseCase.EXPECT().Login(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(&authResp, nil)
		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusOK {
			t.Errorf("Code is %v, wants: %v", resp.Code, http.StatusOK)
		}
	})
	t.Run("Bad request error", func(t *testing.T) {

		loginParam := model.LoginParam{CompanyID: "1", EmployeeNo: "2"}
		param, _ := json.Marshal(loginParam)

		req := httptest.NewRequest("POST", "/auth/login", strings.NewReader(string(param)))
		resp := httptest.NewRecorder()

		eUseCase.EXPECT().Login(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))
		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusBadRequest {
			t.Errorf("Code is %v, wants: %v", resp.Code, http.StatusBadRequest)
		}

	})
	t.Run("Bad request error", func(t *testing.T) {

		loginParam := model.LoginParam{CompanyID: "1", EmployeeNo: "2", Password: "3"}
		param, _ := json.Marshal(loginParam)

		req := httptest.NewRequest("POST", "/auth/login", strings.NewReader(string(param)))
		resp := httptest.NewRecorder()

		eUseCase.EXPECT().Login(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))
		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusInternalServerError {
			t.Errorf("Code is %v, wants: %v", resp.Code, http.StatusInternalServerError)
		}
	})
}

func TestEmployeeHttpHandler_Hello(t *testing.T) {

	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	eUseCase := mock_usecase.NewMockEmployeeUseCase(ctrl)
	eHandler := NewEmployeeHttpHandler(eUseCase)

	router := gin.New()
	router.GET("/employee/hello", eHandler.Hello)

	t.Run("Successful hello", func(t *testing.T) {

		req := httptest.NewRequest("GET", "/employee/hello", nil)
		resp := httptest.NewRecorder()

		eUseCase.EXPECT().Hello(gomock.Any()).Return("", nil)
		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusOK {
			t.Errorf("Code is %v, wants: %v", resp.Code, http.StatusOK)
		}
	})
	t.Run("Server error", func(t *testing.T) {

		req := httptest.NewRequest("GET", "/employee/hello", nil)
		resp := httptest.NewRecorder()

		eUseCase.EXPECT().Hello(gomock.Any()).Return("", errors.New("error"))
		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusInternalServerError {
			t.Errorf("Code is %v, wants: %v", resp.Code, http.StatusInternalServerError)
		}
	})
}

func TestEmployeeHttpHandler_CreateNotice(t *testing.T) {

	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	eUseCase := mock_usecase.NewMockEmployeeUseCase(ctrl)
	eHandler := NewEmployeeHttpHandler(eUseCase)

	router := gin.New()
	router.POST("/employee/notice", eHandler.CreateNotice)

	t.Run("Notice created", func(t *testing.T) {

		noticeParam := model.CreateNoticeParam{CompanyID: "1", EmployeeNo: "2", Type: enum.SICK, Visibility: enum.PRIVATE, PeriodStart: "2019-04-20", PeriodEnd: "2019-04-20"}
		param, _ := json.Marshal(noticeParam)

		req := httptest.NewRequest("POST", "/employee/notice", strings.NewReader(string(param)))
		resp := httptest.NewRecorder()

		eUseCase.EXPECT().CanCreateNotice(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		eUseCase.EXPECT().CreateNotice(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusCreated {
			t.Errorf("Code is %v, wants: %v", resp.Code, http.StatusCreated)
		}
	})
	t.Run("Can not create for other employee", func(t *testing.T) {

		noticeParam := model.CreateNoticeParam{CompanyID: "1", EmployeeNo: "2", Type: enum.SICK, Visibility: enum.PRIVATE, PeriodStart: "2019-04-20", PeriodEnd: "2019-04-20"}
		param, _ := json.Marshal(noticeParam)

		req := httptest.NewRequest("POST", "/employee/notice", strings.NewReader(string(param)))
		resp := httptest.NewRecorder()

		eUseCase.EXPECT().CanCreateNotice(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("error"))
		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusForbidden {
			t.Errorf("Code is %v, wants: %v", resp.Code, http.StatusForbidden)
		}
	})
	t.Run("Server error", func(t *testing.T) {

		noticeParam := model.CreateNoticeParam{CompanyID: "1", EmployeeNo: "2", Type: enum.SICK, Visibility: enum.PRIVATE, PeriodStart: "2019-04-20", PeriodEnd: "2019-04-20"}
		param, _ := json.Marshal(noticeParam)

		req := httptest.NewRequest("POST", "/employee/notice", strings.NewReader(string(param)))
		resp := httptest.NewRecorder()

		eUseCase.EXPECT().CanCreateNotice(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		eUseCase.EXPECT().CreateNotice(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("error"))
		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusInternalServerError {
			t.Errorf("Code is %v, wants: %v", resp.Code, http.StatusInternalServerError)
		}
	})
	t.Run("Invalid periodStart date", func(t *testing.T) {

		noticeParam := model.CreateNoticeParam{CompanyID: "1", EmployeeNo: "2", Type: enum.SICK, Visibility: enum.PRIVATE, PeriodStart: "2019", PeriodEnd: "2019-04-20"}
		param, _ := json.Marshal(noticeParam)

		req := httptest.NewRequest("POST", "/employee/notice", strings.NewReader(string(param)))
		resp := httptest.NewRecorder()

		eUseCase.EXPECT().CanCreateNotice(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		eUseCase.EXPECT().CreateNotice(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("error"))
		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusBadRequest {
			t.Errorf("Code is %v, wants: %v", resp.Code, http.StatusBadRequest)
		}
	})
	t.Run("Invalid periodEnd date", func(t *testing.T) {

		noticeParam := model.CreateNoticeParam{CompanyID: "1", EmployeeNo: "2", Type: enum.SICK, Visibility: enum.PRIVATE, PeriodStart: "2019-04-20", PeriodEnd: "2019"}
		param, _ := json.Marshal(noticeParam)

		req := httptest.NewRequest("POST", "/employee/notice", strings.NewReader(string(param)))
		resp := httptest.NewRecorder()

		eUseCase.EXPECT().CanCreateNotice(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		eUseCase.EXPECT().CreateNotice(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("error"))
		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusBadRequest {
			t.Errorf("Code is %v, wants: %v", resp.Code, http.StatusBadRequest)
		}
	})
}

func TestEmployeeHttpHandler_GetCompanyNotice(t *testing.T) {

	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	eUseCase := mock_usecase.NewMockEmployeeUseCase(ctrl)
	eHandler := NewEmployeeHttpHandler(eUseCase)

	router := gin.New()
	router.GET("/employee/company-notice", eHandler.GetCompanyNotice)

	t.Run("Notice returned", func(t *testing.T) {

		param := url.Values{}
		param.Add("companyId", "1")

		req := httptest.NewRequest("GET", "/employee/company-notice", nil)
		req.URL.RawQuery = param.Encode()
		resp := httptest.NewRecorder()

		eUseCase.EXPECT().GetCompanyNotice(gomock.Any(), gomock.Any()).Return(nil, nil)
		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusOK {
			t.Errorf("Code is %v, wants: %v", resp.Code, http.StatusOK)
		}
	})
	t.Run("Bad request", func(t *testing.T) {

		req := httptest.NewRequest("GET", "/employee/company-notice", nil)
		resp := httptest.NewRecorder()

		eUseCase.EXPECT().GetCompanyNotice(gomock.Any(), gomock.Any()).Return(nil, nil)
		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusBadRequest {
			t.Errorf("Code is %v, wants: %v", resp.Code, http.StatusBadRequest)
		}
	})
}

func TestEmployeeHttpHandler_GetOwnNotice(t *testing.T) {

	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	eUseCase := mock_usecase.NewMockEmployeeUseCase(ctrl)
	eHandler := NewEmployeeHttpHandler(eUseCase)

	router := gin.New()
	router.GET("/employee/notice", eHandler.GetOwnNotice)

	t.Run("Notice returned", func(t *testing.T) {

		req := httptest.NewRequest("GET", "/employee/notice", nil)
		resp := httptest.NewRecorder()

		eUseCase.EXPECT().GetNotice(gomock.Any()).Return(nil, nil)
		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusOK {
			t.Errorf("Code is %v, wants: %v", resp.Code, http.StatusOK)
		}
	})
	t.Run("Server error", func(t *testing.T) {

		req := httptest.NewRequest("GET", "/employee/notice", nil)
		resp := httptest.NewRecorder()

		eUseCase.EXPECT().GetNotice(gomock.Any()).Return(nil, errors.New("error"))
		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusInternalServerError {
			t.Errorf("Code is %v, wants: %v", resp.Code, http.StatusInternalServerError)
		}
	})
}