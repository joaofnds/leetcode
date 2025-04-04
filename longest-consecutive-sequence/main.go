package main

import "sort"

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)

	score, maxScore := 1, 1

	for i := 1; i < len(nums); i++ {
		prev, curr := nums[i-1], nums[i]

		if prev == curr {
			continue
		}

		if curr == prev+1 {
			score++
			continue
		}

		maxScore = max(maxScore, score)
		score = 1
	}

	return max(maxScore, score)
}

func longestConsecutive_On(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	numSet := map[int]bool{}
	for _, num := range nums {
		numSet[num] = true
	}

	maxScore := 1

	for num := range numSet {
		if !numSet[num-1] {
			score := 1

			for numSet[num+1] {
				num++
				score++
			}

			maxScore = max(maxScore, score)
		}
	}

	return maxScore
}
