package main

type car struct {
	name  string
	power int
}

func (c *car) setName(name string) {
	c.name = name
}
func (c *car) getName() string {
	return c.name
}
func (c *car) setPower(power int) {
	c.power = power
}
func (c *car) getPower() int {
	return c.power
}
