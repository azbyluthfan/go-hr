package token

import (
	"github.com/dgrijalva/jwt-go"
	"testing"
	"time"
)

func TestNewJwtGenerator(t *testing.T) {
	duration, _ := time.ParseDuration("1h")
	gen := NewJwtGenerator("secret-key", duration, duration)

	switch gen.(type) {
	case AccessTokenGenerator:
		return
	default:
		t.Errorf("Bad object type returned")
	}
}

func TestJwtGenerator_GenerateAccessToken(t *testing.T) {
	duration, _ := time.ParseDuration("1h")
	gen := NewJwtGenerator("secret-key", duration, duration)

	expired := time.Now().Add(duration)
	claim := jwt.StandardClaims{
		Audience: "www.go.hr",
		ExpiresAt: expired.Unix(),
		Id: "123",
	}

	accessToken := <-gen.GenerateAccessToken(BearerClaims{
		"10000",
		"abc",
		"normal",
		"def",
		true,
		claim,
	})

	if accessToken.AccessToken.ExpiredAt != expired {
		t.Errorf("Incorrect expiredAt value, got: %s, want: %s", accessToken.AccessToken.ExpiredAt, expired)
	}
}