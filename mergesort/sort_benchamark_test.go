package mergesort_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/droxer/algo/mergesort"
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

func BenchmarkMSort100(b *testing.B) {
	benchmarkMSort(100, b)
}

func BenchmarkMSort1000(b *testing.B) {
	benchmarkMSort(1000, b)
}

func BenchmarkMSort10000(b *testing.B) {
	benchmarkMSort(10000, b)
}

func benchmarkMSort(i int, b *testing.B) {
	for j := 0; j < b.N; j++ {
		b.StopTimer()
		values := perm(i)
		b.StartTimer()
		mergesort.Sort(values)
	}
}
