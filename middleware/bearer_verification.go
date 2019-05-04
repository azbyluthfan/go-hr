package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Verify access token and check if claim is valid or not
// Pass parsed claims into context
func BearerVerify(secret string) gin.HandlerFunc {
	return func (c *gin.Context) {
		req, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
			b := ([]byte(secret))
			return b, nil
		})

		// parse claims and set to context
		if req != nil && req.Claims != nil && req.Claims.(jwt.MapClaims) != nil {
			c.Set("employeeNo", req.Claims.(jwt.MapClaims)["eno"].(string))
			c.Set("role", req.Claims.(jwt.MapClaims)["role"].(string))
			c.Set("employeeId", req.Claims.(jwt.MapClaims)["eid"].(string))
			c.Set("companyId", req.Claims.(jwt.MapClaims)["cid"].(string))
		}

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		}
	}
}
