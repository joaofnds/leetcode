package main

func canJump(nums []int) bool {
	var furthest int
	for pos, maxJump := range nums {
		if pos > furthest {
			return false
		}
		furthest = max(pos+maxJump, furthest)
	}
	return true
}
