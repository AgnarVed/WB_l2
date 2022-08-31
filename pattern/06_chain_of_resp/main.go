package main

type defect struct {
	code               string
	registrationDone   bool
	classificationDone bool
	criticalDone       bool
}

func main() {
	crit := &critical{}
	clsfr := &classifier{next: crit}
	reg := &registration{next: clsfr}
	problem := &defect{code: "M007"}
	reg.execute(problem)
}
