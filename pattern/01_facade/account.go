package main

import "fmt"

type Account struct {
	name string
	pin  string
}

func NewAccount(name, pin string) *Account {
	return &Account{
		name: name,
		pin:  pin,
	}
}

var database = map[string]string{"Andrey": "1425", "Valya": "1234"}

func (a *Account) VerifyAccount() bool {
	if pass, ok := database[a.name]; ok {
		if pass == a.pin {
			fmt.Printf("\nHello %s.\n", a.name)
			return true
		} else {
			fmt.Println("Wrong input")
			return false
		}
	}
	return false
}

func (a *Account) CreateAccount() bool {
	database[a.name] = a.pin
	fmt.Printf("Account for %s created.\n", a.name)
	return true
}

func (a *Account) DeleteAccount() {
	delete(database, a.name)
	fmt.Printf("Account for %s deleted.\n", a.name)

}
