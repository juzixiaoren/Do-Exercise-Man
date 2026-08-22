package main

func canCompleteCircuit(gas []int, cost []int) int {
	gasNum := 0
	gasNum += 0
	startPos := 0
	totalGas := 0
	for i := 0; i < len(gas); i++ {
		gasNum += gas[i]
		gasNum -= cost[i]
		totalGas += gas[i] - cost[i]
		if gasNum < 0 {
			startPos = i + 1
			gasNum = 0
		}
	}
	if totalGas < 0 {
		return -1
	} else {
		return startPos
	}
}
