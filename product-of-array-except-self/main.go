package main

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))

	prefix := 1
	for i := range len(nums) {
		result[i] = prefix
		prefix *= nums[i]
	}

	suffix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}
