package recursive_test

import (
	"testing"

	"github.com/droxer/algo/recursive"
)

func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		recursive.Factorial(10)
	}
}

func BenchmarkFactorialIter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		recursive.FactorialIter(10)
	}
}

func BenchmarkFactorialCPS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		recursive.FactorialCPS(10, 1, func(result int) {
			_ = result
		})
	}
}

func BenchmarkFactorialTrampoline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		recursive.FactorialTrampoline(10, 1).Run()
	}
}
