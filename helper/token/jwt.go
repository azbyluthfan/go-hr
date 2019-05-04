package token

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// BearerClaims data structure for claims
type BearerClaims struct {
	EmployeeNo	string	`json:"eno"`
	EmployeeID	string	`json:"eid"`
	Role      	string  `json:"role"`
	CompanyID  	string  `json:"cid"`
	Authorized  bool	`json:"authorized,bool"`
	jwt.StandardClaims
}

// AccessToken data structure
type AccessToken struct {
	AccessToken string
	ExpiredAt   time.Time
}

// AccessTokenResponse data structure
type AccessTokenResponse struct {
	Error       error
	AccessToken AccessToken
}

// jwtGenerator private data structure
type jwtGenerator struct {
	secret         	string
	tokenAge        time.Duration
	refreshTokenAge time.Duration
}

// AccessTokenGenerator interface abstraction
type AccessTokenGenerator interface {
	GenerateAccessToken(cl BearerClaims) <-chan AccessTokenResponse
}

// Generate AccessTokenGenerator object
func NewJwtGenerator(secret string, tokenAge, refreshTokenAge time.Duration) AccessTokenGenerator {
	return &jwtGenerator{
		secret: secret,
		tokenAge: tokenAge,
		refreshTokenAge: refreshTokenAge,
	}
}

// Generate access token from BearerClaims
func (j *jwtGenerator) GenerateAccessToken(cl BearerClaims) <-chan AccessTokenResponse {
	result := make(chan AccessTokenResponse)
	go func() {
		defer close(result)

		now := time.Now()
		age := now.Add(j.tokenAge)

		token := jwt.New(jwt.SigningMethodHS256)
		claims := make(jwt.MapClaims)
		claims["iss"] = cl.Issuer
		claims["aud"] = cl.Audience
		claims["exp"] = age.Unix()
		claims["iat"] = now.Unix()
		claims["sub"] = cl.Subject
		claims["eno"] = cl.EmployeeNo
		claims["eid"] = cl.EmployeeID
		claims["role"] = cl.Role
		claims["cid"] = cl.CompanyID
		claims["authorized"] = cl.Authorized
		token.Claims = claims

		tokenString, err := token.SignedString([]byte(j.secret))
		if err != nil {
			result <- AccessTokenResponse{Error: err}
			return
		}
		result <- AccessTokenResponse{Error: nil, AccessToken: AccessToken{AccessToken: tokenString, ExpiredAt: age}}
	}()

	return result
}
