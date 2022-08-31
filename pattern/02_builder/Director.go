package main

type Director struct {
	builder Builder
}

func NewDirector(b Builder) *Director {
	return &Director{builder: b}
}

func (d *Director) setBuilder(b Builder) {
	d.builder = b
}

func (d *Director) buildHouse() House {
	d.builder.SetDoorType()
	d.builder.SetFloorNumber()
	d.builder.SetRoofType()
	d.builder.SetWallMaterial()
	d.builder.SetWindowType()
	return d.builder.GetHouse()
}
