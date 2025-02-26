package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtEnv struct {
	JwtSecret     string
	JwtExpiration int
}

type JwtService interface {
	GenerateToken(userID string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func NewJWTService(secret string, expiration int) JwtService {
	return &JwtEnv{
		JwtSecret:     secret,
		JwtExpiration: expiration,
	}
}

func (j *JwtEnv) GenerateToken(userID string) (string, error) {
	claims := &jwtCustomClaim{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(j.JwtExpiration)).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.JwtSecret))
}

func (j *JwtEnv) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(j.JwtSecret), nil
	})
}

func DecodeJWT(tokenString string) (map[string]interface{}, error) {
	// Pisahkan token menjadi 3 bagian (header, payload, signature)
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		fmt.Println("Error decoding token:", err)
		return nil, err
	}

	// Ambil claims (payload)
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token format")
}
