package handlers

import (
	"BankingSystem/Core/domain"
	"BankingSystem/Core/service"
	"encoding/json"
	"net/http"
)


var input struct{
		AccountNo string 
		Amount    float64
		Pin       string  
	}

type Handlers struct{
   Service *service.BankingService
}

func(h *Handlers)WithdrawAmount(w http.ResponseWriter,r *http.Request){

	if err:=json.NewDecoder(r.Body).Decode(&input);err!=nil{
         http.Error(w,"Failed to Decode data",http.StatusBadRequest)
	}

	h.Service.Withdraw(input.AccountNo,input.Amount,input.Pin)
}

func(h *Handlers)DepositeAmount(w http.ResponseWriter,r *http.Request){

    if err:=json.NewDecoder(r.Body).Decode(&input);err!=nil{
         http.Error(w,"Failed to Decode data",http.StatusBadRequest)
	}

	h.Service.Deposite(input.AccountNo,input.Amount,input.Pin)
}

func(h *Handlers)TransferAmount(w http.ResponseWriter,r *http.Request){
    var input struct{
		fromAccountNo string
		fromAccountPin string
		ToAccountNo  string
		Amount       float64
	}

	if err:=json.NewDecoder(r.Body).Decode(&input);err!=nil{
         http.Error(w,"Failed to Decode data",http.StatusBadRequest)
	}

	h.Service.Transfer(input.fromAccountNo,input.fromAccountPin,input.ToAccountNo,input.Amount)
}

func(h *Handlers)SetPin(w http.ResponseWriter,r *http.Request){
	var input struct{
		AccountNo string
		OldPin    string
		NewPin    string
	}
	if err:=json.NewDecoder(r.Body).Decode(&input);err!=nil{
         http.Error(w,"Failed to Decode data",http.StatusBadRequest)
	}
	h.Service.SetPin(input.AccountNo,input.OldPin,input.NewPin)
}

func(h *Handlers)CreateAccount(w http.ResponseWriter,r *http.Request){
	var input domain.Customer	
	if err:=json.NewDecoder(r.Body).Decode(&input);err!=nil{
         http.Error(w,"Failed to Decode data",http.StatusBadRequest)
	}

	h.Service.CreateAccount(input)
	
}