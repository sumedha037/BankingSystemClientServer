package ports

import "BankingSystem/Core/domain"

type CustomerRepository interface {
	SaveCustomer(customer domain.Customer)error
}

type AccountRepository interface{
    GetPin(accountNo string)(string,error)
	ChangePin(accountNo string,Pin string)error
  	CreateAccount(customerId string,accountType string)(domain.Account,error)
	GetBalance(accountNo string)(float64,error)
	IncreaseBalance(accountNo string,amount float64)error
	DecreaseBalance(accountNo string,amount float64)error
}

type TransactionRepository interface{
	SaveTransction(transactionId string,fromAccountNo string,toAcountNo string,Amount float64,timestamp string)error
	GetTransactionDetail(transactionId string)(domain.Transaction,error)
}
