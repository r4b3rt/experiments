package defer_test

import "testing"

func sum(max int) int {
	total := 0
	for i := 0; i < max; i++ {
		total += i
	}

	return total
}

func foo() {
	defer func() {
		sum(10)
	}()

	sum(100)
}

func BenchmarkDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		foo()
	}
}
