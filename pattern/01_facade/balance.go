package main

import "fmt"

type Balance struct {
	name  string
	value int
}

func NewBalance(name string) *Balance {
	return &Balance{
		name:  name,
		value: 0,
	}
}

var balanceDB = map[string]int{"Andrey": 100}

func (b *Balance) CheckBalance() (int, bool) {
	if _, ok := balanceDB[b.name]; !ok {
		fmt.Printf("Balance for %s doens't exist.\n", b.name)
		fmt.Printf("Creating balance for %s.\n", b.name)
		b.CreateBalance(b.name)
		return 0, false
	} else {
		fmt.Printf("%s balance is %d", b.name, balanceDB[b.name])
		return balanceDB[b.name], true
	}
}

func (b *Balance) CreateBalance(name string) {
	balanceDB[name] = 0
	fmt.Printf("Balance for %s created.\n", b.name)
}

func (b *Balance) AddToBalance(value int) {
	balanceDB[b.name] += value
	b.value += value
	fmt.Printf("Added %d to %s balance\n", value, b.name)
}

func (b *Balance) GetFromBalance(value int) bool {
	if b.value < value {
		return false
	}
	b.value -= value
	balanceDB[b.name] -= value
	return true
}
