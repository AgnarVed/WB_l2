package main

import "fmt"

type Money struct {
	Account *Account
	Balance *Balance
}

func NewBAN(name, pin string) *Money {
	return &Money{
		Account: NewAccount(name, pin),
		Balance: NewBalance(name),
	}
}

func (m *Money) putMoneyToBAN(value int) bool {
	if !m.Account.VerifyAccount() {
		fmt.Printf("Wrong account input\n")
		return false
	}
	if _, ok := m.Balance.CheckBalance(); !ok {
		m.Balance.AddToBalance(value)
		return false
	}
	m.Balance.AddToBalance(value)
	fmt.Printf("Put %d to %s's BAN", value, m.Balance.name)
	return true
}

func (m *Money) getMoneyFromBAN(value int) bool {
	if !m.Account.VerifyAccount() {
		fmt.Printf("Wrong account input\n")
		return false
	}
	if !m.Balance.GetFromBalance(value) {
		fmt.Printf("Not enough money\n")
		return false
	} else {
		fmt.Printf("Get %d from bankomat\n", value)
		fmt.Printf("Balance is %d\n", m.Balance.value)
	}
	return false
}

func (m *Money) checkBAN() bool {
	if !m.Account.VerifyAccount() {
		fmt.Printf("Wrong account input\n")
		return false
	}
	fmt.Printf("%s's BAN is %d\n\n", m.Account.name, m.Balance.value)
	return true
}
