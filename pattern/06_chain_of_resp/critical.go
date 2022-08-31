package main

import "fmt"

type critical struct {
	next point
}

func (m *critical) execute(d *defect) {
	if d.criticalDone {
		fmt.Println("Defect already handled")
		return
	}
	switch d.code {
	case "B0008":
		fmt.Println("Срочно обратитесь в сервис!!!")
	default:
		fmt.Println("Ошибка обработана")
	}
	d.criticalDone = true
}
func (m *critical) setNext(next point) {
	m.next = next
}
