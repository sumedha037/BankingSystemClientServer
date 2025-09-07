package app

import (
	"BankingSystem/Core/service"
	adaptars "BankingSystem/adaptars/db"
	"BankingSystem/adaptars/handlers"
    "net/http"
    "log"
       "github.com/gorilla/mux"

)

func Start(){

    database := GetInstance()
 
   CustomerRepo:=adaptars.NewCustomer(database)
   AccountRepo:=adaptars.NewAccount(database)
   TransactionRepo:=adaptars.NewTransaction(database)

   BankingService:=service.NewBankingService(AccountRepo,CustomerRepo,TransactionRepo)
   
   h:=handlers.NewHandler(BankingService)
    r := mux.NewRouter()
 
    r.HandleFunc("/deposite",h.DepositeAmount).Methods(http.MethodPost)
    r.HandleFunc("/Withdraw", h.WithdrawAmount).Methods(http.MethodPost)
    r.HandleFunc("/transfer",h.TransferAmount).Methods(http.MethodPost)
    r.HandleFunc("/CreateAccount", h.CreateAccount).Methods(http.MethodPost)
    r.HandleFunc("/SetPin",h.SetPin).Methods(http.MethodPost)
 
    log.Println("Server running on:8080")
    http.ListenAndServe(":8080", r)
}