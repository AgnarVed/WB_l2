package main

type ferrari struct {
	car
}

func newFerrari() iCar {
	return &ferrari{
		car: car{
			name:  "ferrari",
			power: 450,
		},
	}
}
