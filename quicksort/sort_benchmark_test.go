package quicksort_test

import (
	"math/rand"
	"testing"
	"time"

	quicksort "github.com/droxer/algo/quicksort"
)

func init() {
	seed := time.Now().Unix()
	rand.Seed(seed)
}

func perm(n int) (out []int) {
	for _, v := range rand.Perm(n) {
		out = append(out, v)
	}
	return
}

func BenchmarkQSort100(b *testing.B) {
	benchmarkQSort(100, b)
}

func BenchmarkQSort1000(b *testing.B) {
	benchmarkQSort(1000, b)
}

func BenchmarkQSort10000(b *testing.B) {
	benchmarkQSort(10000, b)
}

func benchmarkQSort(i int, b *testing.B) {
	for j := 0; j < b.N; j++ {
		b.StopTimer()
		values := perm(i)
		b.StartTimer()
		quicksort.Sort(values)
	}
}
