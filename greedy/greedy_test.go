package greedy

import "testing"

func TestCanCompleteCircuit(t *testing.T) {
	tests := []struct {
		gas  []int
		cost []int
		want int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}, 3},
		{[]int{2, 3, 4}, []int{3, 4, 3}, -1},
	}
	for _, test := range tests {
		if got := CanCompleteCircuit(test.gas, test.cost); got != test.want {
			t.Errorf("canCompleteCircuit(%v, %v) = %d, want %d", test.gas, test.cost, got, test.want)
		}
	}
}
