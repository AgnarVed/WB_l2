package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("fone after %v", time.Since(start))
}
func or(ch ...<-chan interface{}) <-chan interface{} {
	outCh := make(chan interface{})
	var wg sync.WaitGroup
	wg.Add(1)
	for _, channel := range ch {
		go func(channel <-chan interface{}) {
			for range channel {
			}
			wg.Done()
		}(channel)
	}
	wg.Wait()
	close(outCh)
	return outCh
}
