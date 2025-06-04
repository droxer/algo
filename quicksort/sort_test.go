package quicksort_test

import (
    "github.com/droxer/algo/quicksort"
    "math/rand"
    "reflect"
    "testing"
    "time"
)

func init() {
    seed := time.Now().Unix()
    rand.Seed(seed)
}

func perm(n int) (out []int) {
    for _, v := range rand.Perm(n) {
        out = append(out, v)
    }
    return
}

func TestQSort(t *testing.T) {
    values := []int{9, 1, 20, 3, 6, 7}
    expected := []int{1, 3, 6, 7, 9, 20}

    quicksort.Sort(values)

    if !reflect.DeepEqual(values, expected) {
        t.Fatalf("expected %d, actual is %d", 1, values[0])
    }
}

func BenchmarkQSort100(b *testing.B) {
    benchmarkQSort(100, b)
}

func BenchmarkQSort1000(b *testing.B) {
    benchmarkQSort(1000, b)
}

func BenchmarkQSort10000(b *testing.B) {
    benchmarkQSort(10000, b)
}

func benchmarkQSort(i int, b *testing.B) {
    for j := 0; j < b.N; j++ {
        b.StopTimer()
        values := perm(i)
        b.StartTimer()
        quicksort.Sort(values)
    }
}