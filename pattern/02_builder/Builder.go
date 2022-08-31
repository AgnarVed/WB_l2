package main

type Builder interface {
	SetDoorType()
	SetWindowType()
	SetFloorNumber()
	SetRoofType()
	SetWallMaterial()
	GetHouse() House
}

func getBuilder(builderName string) Builder {
	if builderName == "Simple" {
		return NewSimpleHouse()
	}
	if builderName == "Modern" {
		return NewModernHouse()
	}
	return nil
}
