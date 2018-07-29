package main

import (
	"context"
	"fmt"
	"reflect"
	"time"
)

const nMessages = 6

func fanout(message string, propagate func(context.Context, <-chan string, chan<- string, chan<- string, chan<- string) error) {
	fmt.Printf("%s:\n", message)

	var src = make(chan string)
	var dst1 = make(chan string)
	var dst2 = make(chan string)
	var dst3 = make(chan string)
	var done = make(chan interface{})

	go func() {
		for {
			d, ok := <-dst1
			if !ok {
				done <- nil
				return
			}
			time.Sleep(200 * time.Millisecond)
			fmt.Println("dst1:" + d)
		}
	}()
	go func() {
		for {
			d, ok := <-dst2
			if !ok {
				done <- nil
				return
			}
			time.Sleep(100 * time.Millisecond)
			fmt.Println("dst2:" + d)
		}
	}()
	go func() {
		for {
			d, ok := <-dst3
			if !ok {
				done <- nil
				return
			}
			time.Sleep(50 * time.Millisecond)
			fmt.Println("dst3:" + d)
		}

	}()

	go func() {
		for i := 0; i < nMessages; i++ {
			src <- fmt.Sprintf("message%d", i)
		}
		close(src)
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()

	err := propagate(ctx, src, dst1, dst2, dst3)
	if err != nil {
		panic(err)
	}
	close(dst1)
	close(dst2)
	close(dst3)

	for i := 0; i < 3; i++ {
		<-done
	}
}

func propagateSingle(ctx context.Context, src <-chan string, dst1 chan<- string, dst2 chan<- string, dst3 chan<- string) error {
	for {
		data, ok := <-src
		if !ok {
			return nil
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case dst1 <- data:
		case dst2 <- data:
		case dst3 <- data:
		}
	}
}
func propagateMulti(ctx context.Context, src <-chan string, dst1 chan<- string, dst2 chan<- string, dst3 chan<- string) error {
	for {
		data, ok := <-src
		if !ok {
			return nil
		}

		for i := 0; i < 3; i++ {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case dst1 <- data:
			case dst2 <- data:
			case dst3 <- data:
			}
		}
	}
}
func propagateMultiByLoop(ctx context.Context, src <-chan string, dst1 chan<- string, dst2 chan<- string, dst3 chan<- string) error {
	for {
		data, ok := <-src
		if !ok {
			return nil
		}
		var channels = []chan<- string{dst1, dst2, dst3}

		for _, ch := range channels {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case ch <- data:
			}
		}
	}
}
func propagateMultiByReflection(ctx context.Context, src <-chan string, dst1 chan<- string, dst2 chan<- string, dst3 chan<- string) error {
	for {
		data, ok := <-src
		if !ok {
			return nil
		}
		var channels = []chan<- string{dst1, dst2, dst3}

		var cases = make([]reflect.SelectCase, len(channels))

		for i, ch := range channels {
			cases[i] = reflect.SelectCase{
				Chan: reflect.ValueOf(ch),
				Dir:  reflect.SelectSend,
				Send: reflect.ValueOf(data),
			}
		}
		for len(cases) > 0 {
			chosen, _, _ := reflect.Select(cases)
			cases = append(cases[:chosen], cases[chosen+1:]...)
		}
	}
}

func main() {
	fanout("propagateSingle", propagateSingle)
	fanout("propagateMulti", propagateMulti)
	fanout("propagateMultiByLoop", propagateMultiByLoop)
	fanout("propagateMultiByReflection", propagateMultiByReflection)
}
