package main

func minSubArrayLen(target int, nums []int) int {
	left, sum, minLen := 0, 0, len(nums)+1

	for right, num := range nums {
		sum += num

		for sum >= target {
			minLen = min(minLen, right-left+1)
			sum -= nums[left]
			left++
		}
	}

	if minLen > len(nums) {
		return 0
	}

	return minLen
}
