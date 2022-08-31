package main

import "fmt"

// фабрика
func getCar(carType string) (iCar, error) {
	switch carType {
	case "mazda":
		return newMazda(), nil
	case "ferrari":
		return newFerrari(), nil
	default:
		return nil, fmt.Errorf("Wrong car type passed")
	}
}
func main() {
	mazda, _ := getCar("mazda")
	ferrari, _ := getCar("ferrari")
	printDetails(mazda)
	printDetails(ferrari)
}
func printDetails(g iCar) {
	fmt.Printf("Car: %s\n", g.getName())
	fmt.Printf("Power: %d\n", g.getPower())
}
