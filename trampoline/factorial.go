package trampoline



func FactorialTrampoline(n int, acc int) Trampoline {
	if n == 0 {
		return Done{acc}
	}
	return Call{func() Trampoline {
		return FactorialTrampoline(n-1, acc*n)
	}}
}

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

func FactorialIter(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}
