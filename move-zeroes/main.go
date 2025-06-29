package main

func moveZeroes(nums []int) {
	var k int

	for _, n := range nums {
		if n != 0 {
			nums[k] = n
			k++
		}
	}

	for ; k < len(nums); k++ {
		nums[k] = 0
	}
}
