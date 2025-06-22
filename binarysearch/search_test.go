package binarysearch_test

import (
	"math/rand"
	"testing"

	"github.com/droxer/algo/binarysearch"
	"github.com/droxer/algo/quicksort"
)

var (
	sortedArray10000   []int
	sortedArray100000  []int
	sortedArray1000000 []int
)

func sortedRandomArray(n int) []int {
	sliced := rand.Perm(n * 20)[n : n*2-1]
	quicksort.Sort(sliced)
	return sliced
}

func init() {
	sortedArray10000 = sortedRandomArray(10000)
	sortedArray100000 = sortedRandomArray(100000)
	sortedArray1000000 = sortedRandomArray(1000000)
}

func TestBSearch10000(t *testing.T) {
	testBSearch(sortedArray10000, t)
}

func testBSearch(sorted []int, t *testing.T) {
	expected := rand.Intn(len(sorted) - 1)
	target := sorted[expected]

	found := binarysearch.Search(sorted, target)
	if found != expected {
		t.Fatalf("expected %d, actual is %d", expected, found)
	}
}
