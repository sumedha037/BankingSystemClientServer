package service  

import "BankingSystem/Core/domain"
import "BankingSystem/Core/ports"
import "math/rand"
import "strconv"
import "time"

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

func(b *BankingService) Withdraw(accountno string,amount float64,Pin string)error{

	 b.ValidateUser(accountno,Pin)

	 balance,_:=b.AccountRepo.GetBalance(accountno)
	 if balance>amount{
		b.AccountRepo.DecreaseBalance(accountno,amount)
	 }
	 return nil
}

func(b *BankingService) Deposite(accountno string,amount float64,Pin string)error{

	b.ValidateUser(accountno,Pin)

	if amount>0{
		b.AccountRepo.IncreaseBalance(accountno,amount)
	}
  return nil
}

func(b *BankingService) Transfer(fromAccountNo string,fromAccountPin string,toAcountNo string,Amount float64)error{

	b.ValidateUser(fromAccountNo,fromAccountPin)

	//generate transaction Id
	id:=rand.Intn(100000000)
    s:= strconv.Itoa(id)

	timestamp:=time.Now()
	formattedTime := timestamp.Format("2006-01-02 15:04:05")

	balance,_:=b.AccountRepo.GetBalance(fromAccountNo)
	 if balance>Amount{
		b.AccountRepo.DecreaseBalance(fromAccountNo,Amount)
		b.AccountRepo.IncreaseBalance(toAcountNo,Amount)
	 }

	 b.TransactionRepo.SaveTransction(s,fromAccountNo,toAcountNo,Amount,formattedTime)
	 return nil
}

func(b *BankingService)SetPin(accountNo string,OldPin string, NewPin string)error{

	b.ValidateUser(accountNo,OldPin)

	b.AccountRepo.ChangePin(accountNo,NewPin)

   return nil
}

func(b *BankingService)CreateAccount(customer domain.Customer)domain.Account{

    b.CustomerRepo.SaveCustomer(customer)

    //generate Unique Account Number

	b.AccountRepo.CreateAccount(customer.CustomerId,customer.AccountType)

    return domain.Account{}
}