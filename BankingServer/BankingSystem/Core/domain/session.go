package domain

import "github.com/golang-jwt/jwt/v5"

type Claims struct{
	AccountNo string
	jwt.RegisteredClaims
}