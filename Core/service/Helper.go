package service 
import "log"
import "fmt"
import "os"
import "strconv"



func(b *BankingService) ValidateUser(accountNo string,Pin string)(bool,error){

   s,err:=b.AccountRepo.GetPin(accountNo)
   if err!=nil{
    return false,fmt.Errorf("unable to get the pin %v",err)
   }

   if s!=Pin{
    return false,fmt.Errorf("invalid Password")
   }
   log.Println(s)
 return true,nil
}



func (b *BankingService) IncreaseAmount(accountNo string,amount float64)error{
   currentAmount,err:=b.AccountRepo.GetBalance(accountNo)
  if err!=nil{
    log.Printf("unable to get current balance %v",err)
  }
  currentAmount+=amount
  return b.AccountRepo.SaveBalance(accountNo,currentAmount)
}



func (b *BankingService) DecreaseAmount(accountNo string,amount float64)error{
   currentAmount,err:=b.AccountRepo.GetBalance(accountNo)
  if err!=nil{
    log.Printf("unable to get current balance %v",err)
  }
  if currentAmount<amount{
    return fmt.Errorf("not enough balance")
  }
  currentAmount-=amount
  return b.AccountRepo.SaveBalance(accountNo,currentAmount)
}

func (b *BankingService) GenerateSequentialID(length int) string {

	counter := 0
	data, err := os.ReadFile("counter.txt")
	if err == nil {
		val, convErr := strconv.Atoi(string(data))
		if convErr == nil {
			counter = val
		}
	}


	counter++

	_ = os.WriteFile("counter.txt", []byte(strconv.Itoa(counter)), 0644)

	num := strconv.Itoa(counter)
	for len(num) < length {
		num = "0" + num
	}
	return num
}
