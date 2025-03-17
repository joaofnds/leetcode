package main

func canCompleteCircuit(gas []int, cost []int) int {
	var total, tank, start int
	for i := range gas {
		diff := gas[i] - cost[i]
		total += diff
		tank += diff
		if tank < 0 {
			tank = 0
			start = i + 1
		}
	}

	if total < 0 {
		return -1
	}

	return start
}
