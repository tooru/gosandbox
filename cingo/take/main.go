package main

import "fmt"

func repeat(done <-chan interface{}, values ...interface{}) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			for _, v := range values {
				select {
				case <-done:
					return
				case valueStream <- v:
				}
			}
		}
	}()
	return valueStream
}

func take(done <-chan interface{}, c <-chan interface{}, n int) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)

		for i := 0; i < n; i++ {
			select {
			case <-done:
				return
			case valueStream <- c:
			}
		}
	}()
	return valueStream
}

func main() {
	done := make(chan interface{})
	defer close(done)

	for v := range take(done, repeat(done, 1, 2), 4) {
		fmt.Printf("v=%v\n", v)
	}
}
