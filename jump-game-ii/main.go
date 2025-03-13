package main

func jump(nums []int) int {
	var jumps, currentJumpEnd, farthest int

	for i := range len(nums) - 1 {
		farthest = max(farthest, i+nums[i])

		if i == currentJumpEnd {
			jumps++
			currentJumpEnd = farthest
		}
	}

	return jumps
}
