package middleware

import (
	"BankingSystem/Core/service"
	"context"
	"net/http"
)

type ContextKey string
const AccountKey ContextKey="accountNo"


func AuthMiddleware(next http.Handler)http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token:=r.Header.Get("X-Auth-Token")
		accountNo,err:=service.ValidateJWT(token)
		if err!=nil{
			http.Error(w,"Unauthorized: "+err.Error(),http.StatusUnauthorized)
			return
		}
	    
		ctx:=context.WithValue(r.Context(), AccountKey ,accountNo)
		next.ServeHTTP(w,r.WithContext(ctx))
	})
}