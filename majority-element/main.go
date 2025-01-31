package main

func majorityElement(nums []int) int {
	counts := map[int]int{}

	for _, n := range nums {
		counts[n]++
		if counts[n] > (len(nums) / 2) {
			return n
		}
	}

	panic("No majority element found")
}
