package main

func findDifference(nums1 []int, nums2 []int) [][]int {
	set1, set2 := map[int]bool{}, map[int]bool{}

	for _, num := range nums1 {
		set1[num] = true
	}
	for _, num := range nums2 {
		set2[num] = true
	}

	diff1, diff2 := []int{}, []int{}

	for num := range set1 {
		if _, ok := set2[num]; !ok {
			diff1 = append(diff1, num)
		}
	}
	for num := range set2 {
		if _, ok := set1[num]; !ok {
			diff2 = append(diff2, num)
		}
	}

	return [][]int{diff1, diff2}
}
