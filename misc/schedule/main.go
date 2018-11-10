package main

import (
	"fmt"
	"runtime"
	"sync"
)

const nTasks = 10

func preemptive() {
	var wg sync.WaitGroup
	wg.Add(nTasks)
	for i := 0; i < nTasks; i++ {
		i := i
		go func() {
			defer wg.Done()

			fmt.Printf("frac:start:%d\n", i)
			r := 1
			for j := 2; j < 100*1000*1000; j++ {
				r += j
			}
			fmt.Printf("frac:end  :%d:%d\n", i, r)
		}()
	}
	wg.Wait()
}

func nonPreemptive() {
	var wg sync.WaitGroup
	wg.Add(nTasks)
	for i := 0; i < nTasks; i++ {
		i := i
		go func() {
			defer wg.Done()

			fmt.Printf("frac:start:%d\n", i)
			r := 1
			for j := 2; j < 100*1000*1000; j++ {
				r += j
				if j%(1000*1000) == 0 {
					fmt.Print()
				}
			}
			fmt.Printf("frac:end  :%d:%d\n", i, r)
		}()
	}
	wg.Wait()
}

func main() {
	runtime.GOMAXPROCS(1)

	preemptive()
	fmt.Println()
	nonPreemptive()
}
