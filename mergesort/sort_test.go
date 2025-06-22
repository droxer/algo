package mergesort_test

import (
	"reflect"
	"testing"

	"github.com/droxer/algo/mergesort"
)

func TestMSort(t *testing.T) {
	expected := []int{3, 9, 10, 27, 38, 43, 82}
	orig := []int{38, 27, 43, 3, 9, 82, 10}

	mergesort.Sort(orig)

	if !reflect.DeepEqual(orig, expected) {
		t.Fatalf("expected %v, actual is %v", expected, orig)
	}
}
