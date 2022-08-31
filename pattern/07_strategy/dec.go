package main

type buyStrategy struct {
}

func (bs *buyStrategy) actThis(b *broker) {
	b.bank -= 100
}
