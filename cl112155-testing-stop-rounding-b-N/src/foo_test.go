package main

import (
	"math/rand"
	"testing"
)

func BenchmarkFoo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Intn(100000000)
	}
}
