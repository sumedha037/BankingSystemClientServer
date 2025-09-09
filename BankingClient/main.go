package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const serverURL = "http://localhost:8080"

var jwtToken string 

func main() {
	for {
		if jwtToken == "" {
			mainMenu()
		} else {
			transactionMenu()
		}
	}
}



func mainMenu() {
	fmt.Println("\n==== Banking App ====")
	fmt.Println("1. Create Account")
	fmt.Println("2. Login")
	fmt.Println("3. Set Pin")
	fmt.Println("0. Exit")
	fmt.Print("Enter choice: ")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		createAccount()
	case 2:
		login()
	case 3:
		setPin()
	case 0:
		os.Exit(0)
	default:
		fmt.Println("Invalid choice")
	}
}

func transactionMenu() {
	fmt.Println("\n==== Transactions ====")
	fmt.Println("1. Deposit")
	fmt.Println("2. Withdraw")
	fmt.Println("3. Transfer")
	fmt.Println("4. Check Balance")
	fmt.Println("0. Logout")
	fmt.Print("Enter choice: ")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		deposit()
	case 2:
		withdraw()
	case 3:
		transfer()
	case 4:
		checkBalance()
	case 0:
		jwtToken = "" 
	default:
		fmt.Println("Invalid choice")
	}
}



func createAccount() {
	var Username,Name,Email,Phone,AccountType string
	fmt.Print("Enter Username: ")
	fmt.Scan(&Username)
	fmt.Print("Enter Name: ")
	fmt.Scan(&Name)
	fmt.Print("Enter Email: ")
	fmt.Scan(&Email)
	fmt.Print("Enter PhoneNo: ")
	fmt.Scan(&Phone)
	fmt.Print("Enter Account Type: ")
	fmt.Scan(&AccountType)

	reqBody := map[string]string{
		"CustomerId": Username,
		"Name":  Name,
		"Email":Email,
		"Phone":Phone,
		"AccountType":AccountType,
	}
	data, _ := json.Marshal(reqBody)

	resp, err := http.Post(serverURL+"/CreateAccount", "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var res map[string]string
	json.Unmarshal(body, &res)

	if token, ok := res["token"]; ok {
		jwtToken = token
		fmt.Println("Account created successfully! JWT stored automatically.")
	} else {
		fmt.Println(string(body))
	}
}

func login() {
	var accountNo, pin string
	fmt.Print("Enter account number: ")
	fmt.Scan(&accountNo)
	fmt.Print("Enter pin: ")
	fmt.Scan(&pin)

	reqBody := map[string]string{
		"AccountNo": accountNo,
		"Pin":       pin,
	}
	data, _ := json.Marshal(reqBody)

	resp, err := http.Post(serverURL+"/Login", "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var res map[string]string
	json.Unmarshal(body, &res)

	if token, ok := res["Token"]; ok {
		jwtToken = token
		fmt.Println("Login successful! JWT stored automatically.")
	} else {
		fmt.Println(string(body))
	}
}

func setPin() {
	var accountNo, OldPin,NewPin string
	fmt.Print("Enter account number: ")
	fmt.Scan(&accountNo)
	fmt.Print("Enter current pin: ")
	fmt.Scan(&OldPin)
	fmt.Print("Enter new pin: ")
	fmt.Scan(&NewPin)

	reqBody := map[string]string{
		"AccountNo": accountNo,
		"OldPin":       OldPin,
		"NewPin":       NewPin,
	}
	data, _ := json.Marshal(reqBody)

	resp, err := http.Post(serverURL+"/SetPin", "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func deposit() {
	var amount float64
	var pin string
	fmt.Print("Enter amount: ")
	fmt.Scan(&amount)
	fmt.Print("Enter pin: ")
	fmt.Scan(&pin)

	reqBody := map[string]interface{}{
		"pin":    pin,
		"amount": amount,
	}
	callProtectedAPI("/Deposit", reqBody)
}

func withdraw() {
	var amount float64
	var pin string
	fmt.Print("Enter amount: ")
	fmt.Scan(&amount)
	fmt.Print("Enter pin: ")
	fmt.Scan(&pin)

	reqBody := map[string]interface{}{
		"pin":    pin,
		"amount": amount,
	}
	callProtectedAPI("/Withdraw", reqBody)
}

func transfer() {
	var amount float64
	var pin, toAccount string
	fmt.Print("Enter recipient account number: ")
	fmt.Scan(&toAccount)
	fmt.Print("Enter amount: ")
	fmt.Scan(&amount)
	fmt.Print("Enter pin: ")
	fmt.Scan(&pin)

	reqBody := map[string]interface{}{
		"FromAccountPin": pin,
		"ToAccountNo": toAccount,
		"Amount": amount,
	}
	callProtectedAPI("/Transfer", reqBody)
}

func checkBalance() {
	var pin string
	fmt.Print("Enter pin: ")
	fmt.Scan(&pin)

	reqBody := map[string]interface{}{
		"pin": pin,
	}
	callProtectedAPI("/CheckBalance", reqBody)
}



func callProtectedAPI(endpoint string, reqBody interface{}) {
	data, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", serverURL+endpoint, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth-Token", jwtToken) 

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	if resp.StatusCode!=http.StatusOK{
	    mainMenu()
		return
	}
}