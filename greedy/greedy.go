package greedy

func CanCompleteCircuit(gas []int, cost []int) int {
	totalTank := 0
	currTank := 0
	startIndex := 0

	for i, g := range gas {
		diff := g - cost[i]
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
