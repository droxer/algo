package bubblesort_test

import (
	"testing"

	"github.com/droxer/algo/bubble_sort"
)

func TestBubbleSort(t *testing.T) {
	// Test empty array
	arr := []int{}
	bubblesort.Sort(arr)
	if len(arr) != 0 {
		t.Error("Expected empty array to remain empty")
	}

	// Test single element array
	arr = []int{1}
	bubblesort.Sort(arr)
	if arr[0] != 1 {
		t.Error("Expected single element array to remain unchanged")
	}

	// Test already sorted array
	arr = []int{1, 2, 3, 4, 5}
	bubblesort.Sort(arr)
	expected := []int{1, 2, 3, 4, 5}
	for i := range arr {
		if arr[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected, arr)
			break
		}
	}

	// Test reverse sorted array
	arr = []int{5, 4, 3, 2, 1}
	bubblesort.Sort(arr)
	expected = []int{1, 2, 3, 4, 5}
	for i := range arr {
		if arr[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected, arr)
			break
		}
	}

	// Test array with duplicate elements
	arr = []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	bubblesort.Sort(arr)
	expected = []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}
	for i := range arr {
		if arr[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected, arr)
			break
		}
	}

	// Test array with negative numbers
	arr = []int{-5, 3, -1, 0, 2, -3, 4}
	bubblesort.Sort(arr)
	expected = []int{-5, -3, -1, 0, 2, 3, 4}
	for i := range arr {
		if arr[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected, arr)
			break
		}
	}
}
