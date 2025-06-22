package binarysearch_test

import (
	"math/rand"
	"testing"

	"github.com/droxer/algo/binarysearch"
)

func init() {
	sortedArray10000 = sortedRandomArray(10000)
	sortedArray100000 = sortedRandomArray(100000)
	sortedArray1000000 = sortedRandomArray(1000000)
}

func BenchmarkBSearch10000(b *testing.B) {
	benchmarkBSearch(sortedArray10000, b)
}

func BenchmarkBSearch100000(b *testing.B) {
	benchmarkBSearch(sortedArray100000, b)
}

func BenchmarkBSearch1000000(b *testing.B) {
	benchmarkBSearch(sortedArray1000000, b)
}

func benchmarkBSearch(sorted []int, b *testing.B) {
	for j := 0; j < b.N; j++ {
		target := sorted[rand.Intn(len(sorted)-1)]
		binarysearch.Search(sorted, target)
	}
}
