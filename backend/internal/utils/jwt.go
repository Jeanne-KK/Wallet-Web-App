package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtSecret = []byte("1234")

type Claims struct {
	Mail string `json:"mail`
	jwt.RegisteredClaims
}

//		fucntion generate token
func GenerateToken(mail string) (string, error){
	claims := Claims{
		Mail: mail,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1*time.Hour)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}