package main

func largestAltitude(gain []int) int {
	var highest, current int

	for _, delta := range gain {
		current += delta
		highest = max(highest, current)
	}

	return highest
}
