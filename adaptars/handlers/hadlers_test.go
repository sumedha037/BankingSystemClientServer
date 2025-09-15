package handlers

import (
	"BankingSystem/Core/domain"
	"BankingSystem/Core/service"
	adaptars "BankingSystem/adaptars/db"
	"BankingSystem/middleware"
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckBalance(t *testing.T){

	account:=domain.Account{
		AccountNo: "abc123",
		CustomerId: "cust1",
		AccountType: "Saving Account",
		Balance: 30000,
		Pin: "000123",
	}

	AccountDB:=adaptars.NewAccountDB()
	b:=service.NewBankingService(AccountDB,nil,nil)

	h:=NewHandler(b)

	AccountDB.Account["abc123"]=account

   body := `{"Pin":"000123"}`
   body1 := `{"Pin":"000124"}`
   req:=httptest.NewRequest(http.MethodPost,"/CheckBalance",strings.NewReader(body))
   ctx:=context.WithValue(req.Context(), middleware.AccountKey , "abc123")
   req=req.WithContext(ctx)

   req1:=httptest.NewRequest(http.MethodPost,"/CheckBalance",strings.NewReader(body1))
   ctx1:=context.WithValue(req.Context(), middleware.AccountKey ,"abc123")
	req1=req1.WithContext(ctx1)

   w:=httptest.NewRecorder()
   w1:=httptest.NewRecorder()
   h.CheckBalance(w,req)
   h.CheckBalance(w1,req1)
	resp:=w.Result()
	log.Print(resp)
    resp1:=w1.Result()
	assert.Equal(t,http.StatusOK,resp.StatusCode)
	assert.Equal(t,http.StatusInternalServerError,resp1.StatusCode)
}


func TestWithdraw(t *testing.T){
	account:=domain.Account{
		AccountNo: "abc123",
		CustomerId: "cust1",
		AccountType: "Saving Account",
		Balance: 30000,
		Pin: "000123",
	}

	AccountDB:=adaptars.NewAccountDB()
	b:=service.NewBankingService(AccountDB,nil,nil)

	AccountDB.Account["abc123"]=account

	h:=NewHandler(b)
   
	body := `{"Pin":"000123","Amount":1000}`
	req:=httptest.NewRequest(http.MethodPost,"/Withdraw",strings.NewReader(body))
	 ctx:=context.WithValue(req.Context(), middleware.AccountKey ,"abc123")
   req=req.WithContext(ctx)

	w:=httptest.NewRecorder()

	h.WithdrawAmount(w,req)

	resp:=w.Result()

    assert.Equal(t,http.StatusOK,resp.StatusCode)
}



func TestDepositeAmount(t *testing.T){
	account:=domain.Account{
		AccountNo: "abc123",
		CustomerId: "cust1",
		AccountType: "Saving Account",
		Balance: 30000,
		Pin: "000123",
	}

	AccountDB:=adaptars.NewAccountDB()
	b:=service.NewBankingService(AccountDB,nil,nil)

	AccountDB.Account["abc123"]=account

	h:=NewHandler(b)
   
	body := `{"Pin":"000123","Amount":1000}`
	req:=httptest.NewRequest(http.MethodPost,"/deposite",strings.NewReader(body))
	 ctx:=context.WithValue(req.Context(), middleware.AccountKey ,"abc123")
   req=req.WithContext(ctx)

	w:=httptest.NewRecorder()

	h.DepositeAmount(w,req)

	resp:=w.Result()

    assert.Equal(t,http.StatusOK,resp.StatusCode)
}


func TestTransferAmount(t *testing.T){
	account1:=domain.Account{
		AccountNo: "abc123",
		CustomerId: "cust1",
		AccountType: "Saving Account",
		Balance: 30000,
		Pin: "000123",
	}

	account2:=domain.Account{
		AccountNo: "abc124",
		CustomerId: "cust2",
		AccountType: "Saving Account",
		Balance: 30000,
		Pin: "000567",
	}

	AccountDB:=adaptars.NewAccountDB()
	CustomerDB:=adaptars.NewCustomerDB()
	TransactionDB:=adaptars.NewTransactionDB()
	b:=service.NewBankingService(AccountDB,CustomerDB,TransactionDB)

    AccountDB.Account["abc123"]=account1
	AccountDB.Account["abc124"]=account2

	h:=NewHandler(b)

	body:=`{"FromAccountPin":"000123","ToAccountNo":"abc124","Amount":3000}`

	req:=httptest.NewRequest(http.MethodPost,"/transfer",strings.NewReader(body))
	 ctx:=context.WithValue(req.Context(), middleware.AccountKey ,"abc123")
   req=req.WithContext(ctx)

	w:=httptest.NewRecorder()

	h.TransferAmount(w,req)

	resp:=w.Result()

	assert.Equal(t,http.StatusOK,resp.StatusCode)
}

func TestSetPin(t *testing.T){
	account1:=domain.Account{
		AccountNo: "abc123",
		CustomerId: "cust1",
		AccountType: "Saving Account",
		Balance: 30000,
		Pin: "000123",
	}

	AccountDB:=adaptars.NewAccountDB()
	AccountDB.Account["abc123"]=account1
	b:=service.NewBankingService(AccountDB,nil,nil)
	h:=NewHandler(b)

	body:=`{"AccountNo":"abc123","OldPin":"000123","NewPin":"123456"}`
	req:=httptest.NewRequest(http.MethodPost,"/SetPin",strings.NewReader(body))
	w:=httptest.NewRecorder()

	h.SetPin(w,req)

	resp:=w.Result()

	assert.Equal(t,http.StatusOK,resp.StatusCode)

	if AccountDB.Account["abc123"].Pin!="123456"{
		t.Error("error expexted pin not changed")
	}
}

func TestCreateAccount(t *testing.T){
	customer:=`{
		"CustomerId": "cust1",
		"Name": "Rob",
		"Email": "abc@gmail.com",
		"Phone": "2734682379",
		"AccountType": "Saving Account"
	}`

    AccountDB:=adaptars.NewAccountDB()
	CustomerDB:=adaptars.NewCustomerDB()
	TransactionDB:=adaptars.NewTransactionDB()
	b:=service.NewBankingService(AccountDB,CustomerDB,TransactionDB)

	req:=httptest.NewRequest(http.MethodPost,"/CreateAccount",strings.NewReader(customer))
	w:=httptest.NewRecorder()

	h:=NewHandler(b)
    h.CreateAccount(w,req)
	resp:=w.Result()

	assert.Equal(t,http.StatusOK,resp.StatusCode)
}


func TestAuthHandler(t *testing.T){
account:=domain.Account{
		AccountNo: "abc123",
		CustomerId: "cust1",
		AccountType: "Saving Account",
		Balance: 30000,
		Pin: "000123",
	}

	AccountDB:=adaptars.NewAccountDB()
	b:=service.NewBankingService(AccountDB,nil,nil)

	AccountDB.Account["abc123"]=account

	
	body:=`{"AccountNo":"abc123","Pin":"000123"}`
	req:=httptest.NewRequest(http.MethodPost,"/Login",strings.NewReader(body))
	w:=httptest.NewRecorder()

	h:=NewHandler(b)

	h.Login(w,req)

	resp:=w.Result()

	assert.Equal(t,http.StatusOK,resp.StatusCode)
}

