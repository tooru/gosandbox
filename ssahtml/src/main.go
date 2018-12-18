package main

import "fmt"

func ssa() {
	var w, x, y, z int

	x = 5
	x = x - 3
	if x < 3 {
		y = x * 2
		w = y
	} else {
		y = x - 3
	}
	w = x - y
	z = x + y

	fmt.Printf("%d, %d\n", w, z)
}

func main() {
	ssa()
}
