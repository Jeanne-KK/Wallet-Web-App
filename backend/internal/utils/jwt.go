package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
	"errors"
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

func ValidateToken(tokenString string) (*Claims, error){
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error){
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if ok && token.Valid{
		return claims, nil
	}
	return nil, errors.New("invalid token")
}