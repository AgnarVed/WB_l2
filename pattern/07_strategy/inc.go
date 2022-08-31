package main

type sellStrategy struct {
}

func (bs *sellStrategy) actThis(b *broker) {
	b.bank += 100
}
