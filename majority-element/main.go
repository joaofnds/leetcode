package main

func majorityElement(nums []int) int {
	var maj, count int

	for _, n := range nums {
		if count == 0 {
			maj = n
			count = 1
		} else {
			if maj == n {
				count++
			} else {
				count--
			}
		}
	}

	return maj
}
