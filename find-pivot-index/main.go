package main

func pivotIndex(nums []int) int {
	var sum, leftSum int

	for _, num := range nums {
		sum += num
	}

	for i, num := range nums {
		if leftSum+num == sum-leftSum {
			return i
		}
		leftSum += num
	}

	return -1
}
