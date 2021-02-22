package main

import (
	"testing"
)

func Benchmark(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		EvenOrOdd(100)
	}
}
