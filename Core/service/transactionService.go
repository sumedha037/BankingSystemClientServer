package service

import (
	"BankingSystem/Core/domain"
	"BankingSystem/Core/ports"
	customerrors "BankingSystem/customErrors"
	"fmt"
	"log"
	"time"
)

type BankingService struct{
     AccountRepo ports.AccountRepository
	 CustomerRepo  ports.CustomerRepository
	 TransactionRepo ports.TransactionRepository
}


func NewBankingService(a ports.AccountRepository,c ports.CustomerRepository,t ports.TransactionRepository)*BankingService{
	return &BankingService{AccountRepo: a,CustomerRepo: c,TransactionRepo: t}
}

func(b *BankingService) Withdraw(accountno string,amount float64,Pin string)error{

	ok,err:=b.ValidateUser(accountno,Pin);if !ok {
		return customerrors.NewServiceError("Withdraw: Unauthorized User",err)
	}
	    if amount<=0{
			return customerrors.NewServiceError("WithDraw DecreaseAmount",fmt.Errorf("negative amount"))
		}
		err=b.DecreaseAmount(accountno,amount)
		if err!=nil{
			return customerrors.NewServiceError("WithDraw DecreaseAmount",err)
		}

	 return nil
}

func(b *BankingService) Deposite(accountno string,amount float64,Pin string)error{
	ok,err:=b.ValidateUser(accountno,Pin);if!ok{
		return customerrors.NewServiceError("Withdraw:Unauthorized User",err)
	}
   
	if amount<=0{
		return customerrors.NewServiceError("Deposite Incraese Amount",fmt.Errorf("negative amount"))
	}else{
		err:=b.IncreaseAmount(accountno,amount)
	    if err!=nil{
			return customerrors.NewServiceError("Deposite Incraese Amount",err)
		}
	}
  return nil
}

func(b *BankingService) Transfer(fromAccountNo string,fromAccountPin string,toAcountNo string,Amount float64)(string,error){
    var status string

	if fromAccountNo==toAcountNo{
		return "",customerrors.NewServiceError("Transfer: r",fmt.Errorf("cannot transfer money in same account"))
	}

	ok,err:=b.ValidateUser(fromAccountNo,fromAccountPin);if !ok{
		return "",customerrors.NewServiceError("Transfer:Unauthorized User",err)
	}

	id:=b.GenerateSequentialID(8)

	timestamp:=time.Now()
	formattedTime := timestamp.Format("2006-01-02 15:04:05")

	err=b.DecreaseAmount(fromAccountNo,Amount)
	if err!=nil{
		return "",customerrors.NewServiceError("transfer",err)
	}
	err=b.IncreaseAmount(toAcountNo,Amount)
	if err!=nil{
		return "",customerrors.NewServiceError("transfer",err)
	}

	status="Successfull"

	b.TransactionRepo.SaveTransaction(id,fromAccountNo,toAcountNo,Amount,formattedTime,status)
	 return id,nil
}


func(b *BankingService)SetPin(accountNo string,OldPin string, NewPin string)error{

	ok,err:=b.ValidateUser(accountNo,OldPin);if !ok{
		return customerrors.NewServiceError("SetPin:Unauthorized User",err)
	}
  
	err=b.AccountRepo.ChangePin(accountNo,NewPin)
	if err!=nil{
		return customerrors.NewServiceError("Change Pin",err)
	}
	return nil
}


func(b *BankingService)CreateAccount(customer domain.Customer)domain.Account{
	var account domain.Account
   err:= b.CustomerRepo.SaveCustomer(customer)
   if err!=nil{
	log.Printf("failed to save customer in database %v",err)
	return domain.Account{}
   }

	AccountNo:=b.GenerateSequentialID(12)
	Pin:=b.GenerateSequentialID(6)
	Balance:=0.00

	err=b.AccountRepo.SaveAccount(AccountNo,customer.CustomerId,customer.AccountType,Balance,Pin)
	if err!=nil{
	    log.Printf("Failed to Create an Account %v",err)
		return domain.Account{}
	}

   account,err=b.AccountRepo.GetAccountDetails(AccountNo)
   if err!=nil{
	log.Printf("Failed to Get Account Details for %v",AccountNo)
	return domain.Account{}
   }
    return account
}


func(b *BankingService) Balance(accountno string,Pin string)(float64,error){

	 ok,err:=b.ValidateUser(accountno,Pin);if !ok{
		return 0,customerrors.NewServiceError("Balance Unauthorized User",err)
	 }

	 balance,err:=b.AccountRepo.GetBalance(accountno)
	 if err!=nil{
		log.Println("failed to get the balance")
		return 0,customerrors.NewServiceError("Balance",err)
	 }
	 
	 return balance,nil
}