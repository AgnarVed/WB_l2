package main

// point - интерфейс обработчика
type point interface {
	execute(*defect)
	setNext(point)
}
