package main

func findMaxAverage(nums []int, k int) float64 {
	var maxSum, currentSum int

	for i := range k {
		currentSum += nums[i]
	}
	maxSum = currentSum

	for i := k; i < len(nums); i++ {
		currentSum += nums[i] - nums[i-k]
		maxSum = max(maxSum, currentSum)
	}

	return float64(maxSum) / float64(k)
}
