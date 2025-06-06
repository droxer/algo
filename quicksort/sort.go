package quicksort

// QuickSort is a divide-and-conquer algorithm that recursively breaks down a problem into two or more sub-problems
// until they become simple enough to solve directly. The solutions to the sub-problems are then combined to give
// a solution to the original problem.
//
// Time Complexity: O(n log n)
// Space Complexity: O(log n)
//

func Sort(values []int) {
	sort(values, 0, len(values)-1)
}

func sort(values []int, low int, high int) {
	if low >= high {
		return
	}

	pivot := values[low]
	i := low + 1

	for j := low; j <= high; j++ {
		if pivot > values[j] {
			values[i], values[j] = values[j], values[i]
			i++
		}
	}

	values[low], values[i-1] = values[i-1], pivot

	sort(values, low, i-2)
	sort(values, i, high)
}
