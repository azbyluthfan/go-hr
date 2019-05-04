package main

import (
	"github.com/gin-gonic/gin"
	"os"

	employeePresenter "github.com/azbyluthfan/go-hr/modules/employees/presenter"
)

const (
	defaultPort = "9000"
)

func (s *Service) HttpServe() {
	router := gin.New()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	employeeHandler := employeePresenter.NewEmployeeHttpHandler(s.EmployeeUseCase)

	authGroup := router.Group("/auth")
	employeeHandler.MountAuth(authGroup)

	employeeGroup := router.Group("/employee")
	employeeHandler.Mount(employeeGroup)

	_ = router.Run(":" + port)
}
