package quicksort_test

import (
	"reflect"
	"testing"

	quicksort "github.com/droxer/algo/quicksort"
)


func TestQSort(t *testing.T) {
    values := []int{9, 1, 20, 3, 6, 7}
    expected := []int{1, 3, 6, 7, 9, 20}

    quicksort.Sort(values)

    if !reflect.DeepEqual(values, expected) {
        t.Fatalf("expected %d, actual is %d", 1, values[0])
    }
}
