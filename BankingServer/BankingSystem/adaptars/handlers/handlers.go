package handlers

import (
	"BankingSystem/Core/domain"
	"BankingSystem/Core/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"BankingSystem/middleware"
)



type Handlers struct{
   Service *service.BankingService
}

func NewHandler(s *service.BankingService)*Handlers{
   return &Handlers{Service: s}
}


func(h *Handlers)CheckBalance(w http.ResponseWriter,r *http.Request){
	var input struct{
		Pin       string
	}
	if err:=json.NewDecoder(r.Body).Decode(&input);err!=nil{
     http.Error(w,"Failed to Decode data",http.StatusBadRequest)
		return
	}

	accountNo:=r.Context().Value(middleware.AccountKey).(string)

    balance,err:= h.Service.Balance(accountNo,input.Pin)
	if err!=nil{
	http.Error(w,err.Error(),http.StatusInternalServerError)
		return	
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%.2f", balance)))
}


func(h *Handlers)WithdrawAmount(w http.ResponseWriter,r *http.Request){
   var input struct{
		Amount    float64
		Pin       string  
	}

	if err:=json.NewDecoder(r.Body).Decode(&input);err!=nil{
         http.Error(w,"Failed to Decode data",http.StatusBadRequest)
		 return
	}


	accountNo:=r.Context().Value(middleware.AccountKey).(string)


	err:=h.Service.Withdraw(accountNo,input.Amount,input.Pin)
	if err!=nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Withraw Sucessfull"))
}

func(h *Handlers)DepositeAmount(w http.ResponseWriter,r *http.Request){
   var input struct{
		Amount    float64
		Pin       string  
	}

    if err:=json.NewDecoder(r.Body).Decode(&input);err!=nil{
		log.Println("Decode error:", err)
         http.Error(w,err.Error(),http.StatusBadRequest)
		 return
	}

	accountNo:=r.Context().Value(middleware.AccountKey).(string)

	err:=h.Service.Deposite(accountNo,input.Amount,input.Pin)
	if err!=nil{
		  http.Error(w,err.Error(),http.StatusBadRequest)
		  return
	}
	
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Deposite Successfull"))

}

func(h *Handlers)TransferAmount(w http.ResponseWriter,r *http.Request){
    var input struct{
		FromAccountPin string
		ToAccountNo  string
		Amount       float64
	}

	if err:=json.NewDecoder(r.Body).Decode(&input);err!=nil{
         http.Error(w,"Failed to Decode data",http.StatusBadRequest)
		 return
	}

	accountNo:=r.Context().Value(middleware.AccountKey).(string)

	s,err:=h.Service.Transfer(accountNo,input.FromAccountPin,input.ToAccountNo,input.Amount)
    if err!=nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s))
}


func(h *Handlers)SetPin(w http.ResponseWriter,r *http.Request){
	var input struct{
		AccountNo   string
		OldPin    string
		NewPin    string
	}
	if err:=json.NewDecoder(r.Body).Decode(&input);err!=nil{
         http.Error(w,"Failed to Decode data",http.StatusBadRequest)
		 return
	}

	err:=h.Service.SetPin(input.AccountNo,input.OldPin,input.NewPin)
	if err!=nil{
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pin changed successfully"))
}


func(h *Handlers)CreateAccount(w http.ResponseWriter,r *http.Request){
	var input domain.Customer	
	if err:=json.NewDecoder(r.Body).Decode(&input);err!=nil{
         http.Error(w,"Failed to Decode data",http.StatusBadRequest)
		 return
	}

    Account:=h.Service.CreateAccount(input)
	
	w.Header().Set("content-type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Account)
}