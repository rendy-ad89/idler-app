package util

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret []byte = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	ID int64 `json:"id"`
	jwt.MapClaims
}

// GenerateToken generate tokens used for auth
func GenerateToken(id int64) (string, error) {
	claims := Claims{
		id,
		jwt.MapClaims{
			"id":  id,
			"exp": time.Now().Add(time.Hour * 1).Unix(),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}

// ParseToken parsing token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
