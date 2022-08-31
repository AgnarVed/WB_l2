package main

type broker struct {
	bank     int
	strategy iStrategy
}

func initBroker(s iStrategy) *broker {
	return &broker{
		bank:     1000,
		strategy: s,
	}
}
func (b *broker) setStrategy(s iStrategy) {
	b.strategy = s
}
func (b *broker) actThis() {
	b.strategy.actThis(b)
}
