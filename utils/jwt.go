package utils

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	Data string `json:"username"`
	jwt.StandardClaims
}
