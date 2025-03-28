package main

func insert(intervals [][]int, newInterval []int) [][]int {
	result := make([][]int, 0, len(intervals)+1)

	for _, interval := range intervals {
		if interval[1] < newInterval[0] {
			result = append(result, interval)
		} else if interval[0] > newInterval[1] {
			result = append(result, newInterval)
			newInterval = interval
		} else {
			newInterval[0] = min(newInterval[0], interval[0])
			newInterval[1] = max(newInterval[1], interval[1])
		}
	}

	result = append(result, newInterval)

	return result
}
