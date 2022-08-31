package main

type iStrategy interface {
	actThis(b *broker)
}
