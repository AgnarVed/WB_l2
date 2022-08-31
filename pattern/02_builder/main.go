package main

import "fmt"

func main() {
	simpleBuilder := getBuilder("Simple")
	//modernBuilder := getBuilder("Modern")

	director := NewDirector(simpleBuilder)
	simpleHouse := director.buildHouse()

	fmt.Println(simpleHouse.WallMaterial)
	fmt.Println(simpleHouse.FloorNumber)
	fmt.Println(simpleHouse.DoorType)

}
