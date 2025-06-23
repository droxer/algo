package trampoline_test

import (
	"testing"

	"github.com/droxer/algo/trampoline"
)

func BenchmarkFactorialTrampoline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		trampoline.FactorialTrampoline(10, 1).Run()
	}
}

func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		trampoline.Factorial(10)
	}
}

func BenchmarkFactorialIter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		trampoline.FactorialIter(10)
	}
}