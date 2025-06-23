package miniheap_test

import (
	"container/heap"
	"testing"

	"github.com/droxer/algo/miniheap"
)

func TestMinHeap(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{3, 1, 2}, []int{1, 2, 3}},
	}
	for _, test := range tests {
		minHeap := &miniheap.MinHeap{}
		for _, num := range test.nums {
			heap.Push(minHeap, num)
		}
		for i := 0; i < len(test.nums); i++ {
			popped := heap.Pop(minHeap).(int)
			if popped != test.want[i] {
				t.Errorf("popped[%d] = %d, want %d", i, popped, test.want[i])
			}
		}
	}
}
