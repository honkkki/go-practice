package testing

import "testing"

func BenchmarkFib1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib1(30)
	}
}

func BenchmarkFib2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib2(30)
	}
}

func BenchmarkFib3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib3(30)
	}
}
