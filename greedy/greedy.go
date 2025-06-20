package greedy

func canCompleteCircuit(gas []int, cost []int) int {
	totalGas := 0
	totalCost := 0
	for i := 0; i < len(gas); i++ {
		totalGas += gas[i]
		totalCost += cost[i]
	}
	if totalGas < totalCost {
		return -1
	}

	start := 0
	currentGas := 0
	for i := 0; i < len(gas); i++ {
		currentGas += gas[i] - cost[i]
		if currentGas < 0 {
			start = i + 1
			currentGas = 0
		}
	}
	return start
}