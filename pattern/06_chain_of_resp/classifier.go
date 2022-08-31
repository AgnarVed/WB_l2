package main

import "fmt"

// ещё один обработчик
type classifier struct {
	next point
}

func (c *classifier) execute(d *defect) {
	if d.classificationDone {
		fmt.Println("classifier checkup already done")
		c.next.execute(d)
		return
	}
	var v string
	switch d.code {
	case "1368":
		v = "Ошибка стартера"
	case "B0008":
		v = "Пропуски зажигания"
	case "M12":
		v = "Низкое напряжение АКБ"
	default:
		v = "Неизвестная ошибка"
	}
	fmt.Printf("Найдена следующая неисправность: %s\n", v)
	d.classificationDone = true
	c.next.execute(d)
}
func (c *classifier) setNext(next point) {
	c.next = next
}
