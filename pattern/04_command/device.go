package main

import "fmt"

// интерфейс для объекта, который получает команды
type device interface {
	on()
	off()
}

// конкретная реализация получателя
type car struct {
	isRunning bool
}

func (t *car) on() {
	t.isRunning = true
	fmt.Println("Turning car on")
}
func (t *car) off() {
	t.isRunning = false
	fmt.Println("Turning car off")
}
