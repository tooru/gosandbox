package main

import (
	"fmt"
	"time"
)

func main() {
	oct := time.Date(2018, 10, 31, 0, 0, 0, 0, time.UTC)

	fmt.Println(oct.AddDate(0, 1, 0).AddDate(0, 1, 0))
	fmt.Println(oct.AddDate(0, 2, 0))
}
