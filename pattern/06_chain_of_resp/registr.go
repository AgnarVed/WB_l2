package main

import "fmt"

// обработчик 1
type registration struct {
	next point
}

func (r *registration) execute(d *defect) {
	if d.registrationDone {
		fmt.Println("crash registration already done")
		r.next.execute(d)
		return
	}
	fmt.Printf("was detected %s crash", d.code)
	d.registrationDone = true
	r.next.execute(d)
}
func (r *registration) setNext(next point) {
	r.next = next
}
