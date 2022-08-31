package main

type command interface {
	execute()
}

// конкретная команда 1
type powerOnCommand struct {
	device device
}

// реализация
func (c *powerOnCommand) execute() {
	c.device.on()
}

// конкретная команда 2
type powerOffCommand struct {
	device device
}

// реализация
func (c *powerOffCommand) execute() {
	c.device.off()
}
