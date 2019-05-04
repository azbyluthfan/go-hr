package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBearerVerify(t *testing.T) {

	gin.SetMode(gin.TestMode)

	router := gin.New()

	secretKey := "this-is-secret"

	group:= router.Group("/")
	group.Use(BearerVerify(secretKey))
	{
		group.GET("/authenticated", func(context *gin.Context) {})
	}

	t.Run("Valid access token is passed", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/authenticated", nil)
		req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJ3d3cuZ28uaHIiLCJhdXRob3JpemVkIjp0cnVlLCJjaWQiOiJhMGI4ZGM3My02N2U5LTExZTktOTFmNC0wMjQyYWMxMjAwMDIiLCJlaWQiOiJmNWMxZWIxMy02OTdlLTExZTktYTMxZC0wMjQyYWMxMjAwMDMiLCJlbm8iOiIxMDAwMCIsImV4cCI6MTU1Njk4Mzc1MywiaWF0IjoxNTU2OTgwMTUzLCJpc3MiOiIiLCJyb2xlIjoiYWRtaW4iLCJzdWIiOiIifQ.s7t30JukMhMJLQi4BPOUYVj0REy54E5GBdfaf8Vv_6Y")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusOK {
			t.Errorf("Code is %v, wants: %v", resp.Code, http.StatusOK)
		}
	})

	t.Run("Valid access token is passed", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/authenticated", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusUnauthorized {
			t.Errorf("Code is %v, wants: %v", resp.Code, http.StatusUnauthorized)
		}
	})


}