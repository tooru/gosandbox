package main

import (
	"context"
	"fmt"
	"reflect"
	"time"
)

const nMessages = 6

func fanout(message string, dsts [3]chan string, propagate func(context.Context, <-chan string, [3]chan<- string) error) {
	fmt.Printf("%s:\n", message)

	var src = make(chan string)
	var dst1 = dsts[0]
	var dst2 = dsts[1]
	var dst3 = dsts[2]
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

	err := propagate(ctx, src, [3]chan<- string{dst1, dst2, dst3})
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

func propagateSingle(ctx context.Context, src <-chan string, dsts [3]chan<- string) error {
	for {
		data, ok := <-src
		if !ok {
			return nil
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case dsts[0] <- data:
		case dsts[1] <- data:
		case dsts[2] <- data:
		}
	}
}
func propagateMulti(ctx context.Context, src <-chan string, dsts [3]chan<- string) error {
	for {
		data, ok := <-src
		if !ok {
			return nil
		}

		for i := 0; i < 3; i++ {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case dsts[0] <- data:
			case dsts[1] <- data:
			case dsts[2] <- data:
			}
		}
	}
}
func propagateMultiByLoop(ctx context.Context, src <-chan string, dsts [3]chan<- string) error {
	for {
		data, ok := <-src
		if !ok {
			return nil
		}

		for _, ch := range dsts {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case ch <- data:
			}
		}
	}
}
func propagateMultiByReflection(ctx context.Context, src <-chan string, dsts [3]chan<- string) error {
	for {
		data, ok := <-src
		if !ok {
			return nil
		}

		var cases = make([]reflect.SelectCase, len(dsts))

		for i, ch := range dsts {
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

func newChannels(cap int) [3]chan string {
	var res [3]chan string
	for i := 0; i < len(res); i++ {
		res[i] = make(chan string, cap)
	}
	return res
}

func main() {
	fanout("propagateSingle", newChannels(0), propagateSingle)
	fanout("propagateMulti", newChannels(0), propagateMulti)
	fanout("propagateMultiByLoop", newChannels(0), propagateMultiByLoop)
	fanout("propagateMultiByReflection", newChannels(0), propagateMultiByReflection)
	fanout("propagateMultiByLoopWithCapacity", newChannels(5), propagateMultiByLoop)
}
