package main  

import "fmt"

func main(){
    var input int

	fmt.Println("Type 1 if you already have an account")
	fmt.Println("type 2 if you want to create an account")
    
	fmt.Scan(&input)


}


func ChooseState(input int){
	if input==1{
		fmt.Println("What Operation you want to Perform")
		
	}
}