package dp

import "testing"

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 1},
		{2, 2},
		{3, 4},
		{4, 7},
		{5, 13},
		{6, 24},
		{7, 44},
		{8, 81},
		{9, 149},
		{10, 274},
	}
	for _, test := range tests {
		if got := climbStairs(test.n); got != test.want {
			t.Errorf("climbStairs(%d) = %d, want %d", test.n, got, test.want)
		}
	}
}
