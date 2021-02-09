package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// secret key
var secretKey = []byte("abcd1234!@#$")

type ArithmeticCustomClaims struct {
	UserId string `json:"userId"`
	Name   string `json:"name"`

	jwt.StandardClaims
}

func JwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return secretKey, nil
}

func Sign(name, uid string) (string, error) {
	expAt := time.Now().Add(time.Duration(2) * time.Minute).Unix()

	claims := ArithmeticCustomClaims{
		UserId: uid,
		Name:   name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expAt,
			Issuer:    "system",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}
