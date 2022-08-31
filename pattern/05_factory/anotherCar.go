package main

type mazda struct {
	car
}

func newMazda() iCar {
	return &mazda{
		car: car{
			name:  "mazda",
			power: 144,
		},
	}
}
