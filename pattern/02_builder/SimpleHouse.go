package main

type SimpleHouse struct {
	DoorType     string
	WindowType   string
	FloorNumber  int
	RoofType     string
	WallMaterial string
}

func (s *SimpleHouse) SetDoorType() {
	s.DoorType = "Wooden"
}

func (s *SimpleHouse) SetWindowType() {
	s.WindowType = "Wooden with simple glass"
}

func (s *SimpleHouse) SetFloorNumber() {
	s.FloorNumber = 1
}

func (s *SimpleHouse) SetRoofType() {
	s.RoofType = "Triangle"
}

func (s *SimpleHouse) SetWallMaterial() {
	s.WallMaterial = "Wood"
}

func (s *SimpleHouse) GetHouse() House {
	return House{
		DoorType:     s.DoorType,
		WindowType:   s.WindowType,
		FloorNumber:  s.FloorNumber,
		RoofType:     s.RoofType,
		WallMaterial: s.WallMaterial,
	}
}

func NewSimpleHouse() *SimpleHouse {
	return &SimpleHouse{}
}
