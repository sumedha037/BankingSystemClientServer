package adaptars

import "BankingSystem/Core/domain"


type CustomerDB struct{}

func (c *CustomerDB)SaveCustomer(customer domain.Customer)error{return nil}



type AccountDB struct {}

func(d *AccountDB) GetPin(accountNo string)(string,error){return "",nil }

func(d *AccountDB)ChangePin(accountNo string,Pin string)error{return nil}

func(d *AccountDB)CreateAccount(AccountNo string,customerId string,accountType string,Balance float64,Pin string)(domain.Account,error){return domain.Account{},nil}

func(d *AccountDB)GetBalance(accountNo string)(float64,error){return 0,nil}

func(d *AccountDB)IncreaseBalance(accountNo string,amount float64)error{return nil}
	
func(d *AccountDB)DecreaseBalance(accountNO string,amount float64)error{return nil}



type TransactionDB struct{}

func(t *TransactionDB)SaveTransction(transactionId string,fromAccountNo string,toAcountNo string,Amount float64,time string)error{return nil}

func(t *TransactionDB)GetTransactionDetail(transactionId string)(domain.Transaction,error){return domain.Transaction{},nil}