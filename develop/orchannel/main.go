package main

import (
	"fmt"
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
		sig(5*time.Hour),
		sig(7*time.Minute),
	)

	fmt.Printf("done after % v", time.Since(start))
}

// like fan-in patter
func or(channels ...<-chan interface{}) <-chan interface{} {
	or := make(chan interface{})

	output := func(c <-chan interface{}) {
		for n := range c {
			or <- n
		}
		close(or)
	}

	go func() {
		// распаралелливаем
		for _, v := range channels {
			go output(v)
		}
	}()

	return or
}
