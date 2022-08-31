package main

func main() {
	s := &sellStrategy{}
	b := initBroker(s)
	b.actThis()
	b.setStrategy(&buyStrategy{})
	b.actThis()
}
