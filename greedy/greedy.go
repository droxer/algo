package greedy

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	totalTank := 0
	currTank := 0
	startIndex := 0

	for i := 0; i < n; i++ {
		diff := gas[i] - cost[i]
		totalTank += diff
		currTank += diff

		if currTank < 0 {
			startIndex = i + 1
			currTank = 0
		}
	}

	if totalTank < 0 {
		return -1
	}

	return startIndex
}
