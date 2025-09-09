package service

import (
	"BankingSystem/Core/domain"
	"errors"

	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret=[]byte("BeSafe")

func GenerateJWT(accountNo string)(string,error){
	claims:=&domain.Claims{
		AccountNo: accountNo,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10*time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString(jwtSecret)
}


func ValidateJWT(tokenString string)(string,error){
  token,err:=jwt.ParseWithClaims(tokenString,&domain.Claims{},func(token *jwt.Token) (interface{}, error) {
	if _, ok:=token.Method.(*jwt.SigningMethodHMAC);!ok{
	 return nil,errors.New("signing Method Changed")
	}
	return jwtSecret,nil
  })
  if err!=nil{
	return "",err
  }
  if claims,ok:=token.Claims.(*domain.Claims);ok && token.Valid{
	return claims.AccountNo,nil
  }
  return "",errors.New("invalid or expired Token")
}
