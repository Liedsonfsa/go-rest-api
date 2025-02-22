package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = ""

func GenerateToken(email string, userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"userId": userId,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(SecretKey))
}