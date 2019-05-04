package presenter

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"

	"github.com/azbyluthfan/go-hr/middleware"
	"github.com/azbyluthfan/go-hr/modules/employees/model"
	"github.com/azbyluthfan/go-hr/modules/employees/usecase"
)

const (
	secretKey = "C731EFC3E763BD5C6F7E88EB6A883"
)

type employeeHttpHandler struct {
	employeeUseCase usecase.EmployeeUseCase
}

func NewEmployeeHttpHandler(employeeUseCase usecase.EmployeeUseCase) *employeeHttpHandler {
	return &employeeHttpHandler{
		employeeUseCase,
	}
}

func (h *employeeHttpHandler) Mount(group *gin.RouterGroup) {
	group.Use(middleware.BearerVerify(secretKey))
	{
		group.GET("/hello", h.Hello)
		group.GET("/notice", h.GetOwnNotice)
		group.GET("/company-notice", h.GetCompanyNotice)
		group.POST("/notice", h.CreateNotice)
	}
}

func (h *employeeHttpHandler) MountAuth(group *gin.RouterGroup) {
	group.POST("/login", h.Login)
}

func (h *employeeHttpHandler) Login(c *gin.Context) {

	var param model.LoginParam
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authResponse, err := h.employeeUseCase.Login(c, param.CompanyID, param.EmployeeNo, param.Password, secretKey)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, authResponse)
}

func (h *employeeHttpHandler) Hello(c *gin.Context) {
	msg, err := h.employeeUseCase.Hello(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": msg})
}

func (h *employeeHttpHandler) CreateNotice(c *gin.Context) {

	var param model.CreateNoticeParam
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	param.PeriodStart += "T00:00:00Z"
	periodStart, err := time.ParseInLocation(time.RFC3339, param.PeriodStart, time.UTC)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid date for periodStart parameter"})
		return
	}
	param.PeriodEnd += "T00:00:00Z"
	periodEnd, err := time.ParseInLocation(time.RFC3339, param.PeriodEnd, time.UTC)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid date for periodEnd parameter"})
		return
	}

	// check if auth user can create notice for selected employee
	if err = h.employeeUseCase.CanCreateNotice(c, param.CompanyID, param.EmployeeNo); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": err.Error()})
		return
	}

	// create notice
	if err = h.employeeUseCase.CreateNotice(c, param.CompanyID, param.EmployeeNo, param.Type, param.Visibility, periodStart, periodEnd); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Notice created"})
}

func (h *employeeHttpHandler) GetCompanyNotice(c *gin.Context) {

	companyId := c.Query("companyId")
	if companyId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Missing companyId"})
		return
	}

	notices, err := h.employeeUseCase.GetCompanyNotice(c, companyId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": notices})
}

func (h *employeeHttpHandler) GetOwnNotice(c *gin.Context) {

	notices, err := h.employeeUseCase.GetNotice(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": notices})
}