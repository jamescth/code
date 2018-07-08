package util

import "testing"

var lookup [100]int

func fib(n int) int {
	if lookup[n] == 0 {
		if n <= 1 {
			lookup[n] = n
		} else {
			lookup[n] = fib(n-1) + fib(n-2)
		}
	}

	return lookup[n]
}

func fibSlow(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return fibSlow(n-1) + fibSlow(n-2)
}

func TestDpFibFast(t *testing.T) {
	t.Log(fib(40))
}

func TestDpFibSlow(t *testing.T) {
	t.Log(fibSlow(40))
}

func BenchmarkFibFast(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fib(40)
	}
}

func BenchmarkFibSlow(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fibSlow(40)
	}
}
