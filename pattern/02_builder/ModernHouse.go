package main

type ModernHouse struct {
	DoorType     string
	WindowType   string
	FloorNumber  int
	RoofType     string
	WallMaterial string
}

func (m *ModernHouse) SetDoorType() {
	m.DoorType = "Metal"
}

func (m *ModernHouse) SetWindowType() {
	m.WindowType = "Plastic with double glass"
}

func (m *ModernHouse) SetFloorNumber() {
	m.FloorNumber = 9
}

func (m *ModernHouse) SetRoofType() {
	m.RoofType = "Flat"
}

func (m *ModernHouse) SetWallMaterial() {
	m.WallMaterial = "Concrete"
}

func (m *ModernHouse) GetHouse() House {
	return House{
		DoorType:     m.DoorType,
		WindowType:   m.WindowType,
		FloorNumber:  m.FloorNumber,
		RoofType:     m.RoofType,
		WallMaterial: m.WallMaterial,
	}
}

func NewModernHouse() *ModernHouse {
	return &ModernHouse{}
}
