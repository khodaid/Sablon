package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type csrfEnv struct {
	secret  string
	expired int
}

type CsrfService interface {
	GenerateToken() (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

func NewServiceCsrfToken(secret string, expired int) *csrfEnv {
	return &csrfEnv{secret: secret, expired: expired}
}

func (csrf *csrfEnv) GenerateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Second * time.Duration(csrf.expired)).Unix(),
	})

	return token.SignedString([]byte(csrf.secret))
}

func (csrf *csrfEnv) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(csrf.secret), nil
	})
}
