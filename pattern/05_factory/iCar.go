package main

type iCar interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}
