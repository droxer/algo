package bloomfilter_test

import (
	"strconv"
	"testing"

	"github.com/droxer/algo/bloomfilter"
)

func TestBloomFilter(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{3, 1, 2}, []int{1, 2, 3}},
	}
	for _, test := range tests {
		bf := bloomfilter.NewBloomFilter(1000, 3)
		for _, num := range test.nums {
			bf.Add(strconv.Itoa(num))
		}
		for _, num := range test.nums {
			if !bf.Exists(strconv.Itoa(num)) {
				t.Errorf("num %d not found", num)
			}
		}
	}
}