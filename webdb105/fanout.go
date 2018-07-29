package main

import (
	"context"
	"fmt"
)

const nMessages = 6

func main() {
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

	err := call(ctx, src, dst1, dst2, dst3)
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

func call(ctx context.Context, src chan string, dst1 chan string, dst2 chan string, dst3 chan string) error {
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
