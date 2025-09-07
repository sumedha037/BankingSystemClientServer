package service

import (
	"BankingSystem/Core/domain"
	"BankingSystem/Core/ports"
	"fmt"
	"time"
	"log"
)

// type TransactionService interface{
//   WithDraw(accountNo string,Amount float64)error
// 	 Deposite(accountNo string,Amount float64)error
// 	 Transfer(fromAccount domain.Account,toAcountNo string,Amount float64)(domain.Transaction,error)
// 	 SetPin(accountNo string,OldPin string,NewPin string)error
// }



type BankingService struct{
     AccountRepo ports.AccountRepository
	 CustomerRepo  ports.CustomerRepository
	 TransactionRepo ports.TransactionRepository
}


func NewBankingService(a ports.AccountRepository,c ports.CustomerRepository,t ports.TransactionRepository)*BankingService{
	return &BankingService{AccountRepo: a,CustomerRepo: c,TransactionRepo: t}
}

func(b *BankingService) Withdraw(accountno string,amount float64,Pin string)error{

	ok,err:=b.ValidateUser(accountno,Pin);if!ok{
		return err
	}

	 balance,_:=b.AccountRepo.GetBalance(accountno)
	 if balance>amount{
		b.DecreaseAmount(accountno,amount)
	 }
	 return nil
}

func(b *BankingService) Deposite(accountno string,amount float64,Pin string)error{
	ok,err:=b.ValidateUser(accountno,Pin);if!ok{
		return err
	}
	if amount>0{
		err:=b.IncreaseAmount(accountno,amount)
	    if err!=nil{
			return err
		}
	}else{
		return fmt.Errorf("amount is less than zero")
	}
  return nil
}

func(b *BankingService) Transfer(fromAccountNo string,fromAccountPin string,toAcountNo string,Amount float64)(string,error){
    var status string

	ok,err:=b.ValidateUser(fromAccountNo,fromAccountPin);if !ok{
		return "Invalid Pin",err
	}


	id:=b.GenerateSequentialID(8)

	timestamp:=time.Now()
	formattedTime := timestamp.Format("2006-01-02 15:04:05")

	err=b.DecreaseAmount(fromAccountNo,Amount)
	if err!=nil{
		status="Failed"
	}
	err=b.IncreaseAmount(toAcountNo,Amount)
	if err!=nil{
		status="Failed"
	}

	status="Successfull"

	b.TransactionRepo.SaveTransction(id,fromAccountNo,toAcountNo,Amount,formattedTime,status)
	 return id,nil
}

func(b *BankingService)SetPin(accountNo string,OldPin string, NewPin string)error{

	ok,err:=b.ValidateUser(accountNo,OldPin);if !ok{
		return err
	}
  
	err=b.AccountRepo.ChangePin(accountNo,NewPin)
	if err!=nil{
		return err
	}
	return nil
}

func(b *BankingService)CreateAccount(customer domain.Customer)domain.Account{
	var account domain.Account
   err:= b.CustomerRepo.SaveCustomer(customer)
   if err!=nil{
	log.Println("failed to save customer in database")
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
		log.Println("Unauthorized User")
		return 0,err
	 }

	 balance,err:=b.AccountRepo.GetBalance(accountno)
	 if err!=nil{
		log.Println("failed to get the balance")
		return 0,err
	 }
	 
	 return balance,nil
}