package backtracking

import (
	"reflect"
	"testing"
)

func TestSubsets(t *testing.T) {
	tests := []struct {
		nums []int
		want [][]int
	}{
		{[]int{1, 2, 3}, [][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}}},
	}
	for _, test := range tests {
		if got := Subsets(test.nums); !reflect.DeepEqual(got, test.want) {
			t.Errorf("subsets(%v) = %v, want %v", test.nums, got, test.want)
		}
	}
}

func TestPermute(t *testing.T) {
	tests := []struct {
		nums []int
		want [][]int
	}{
		{[]int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
	}
	for _, test := range tests {
		if got := Permute(test.nums); !reflect.DeepEqual(got, test.want) {
			t.Errorf("permute(%v) = %v, want %v", test.nums, got, test.want)
		}
	}
}
