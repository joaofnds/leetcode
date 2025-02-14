package main

func containsNearbyDuplicate(nums []int, k int) bool {
	memo := make(map[int]int, len(nums))

	for i, n := range nums {
		if j, ok := memo[n]; ok && i-j <= k {
			return true
		}

		memo[n] = i
	}

	return false
}
