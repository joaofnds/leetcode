package main

import "strconv"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	i, j := nums[0], nums[0]
	ranges := []string{}

	for _, n := range nums {
		if n-j > 1 {
			if i == j {
				ranges = append(ranges, strconv.Itoa(i))
			} else {
				ranges = append(ranges, strconv.Itoa(i)+"->"+strconv.Itoa(j))
			}
			i = n
		}
		j = n
	}

	if i == j {
		ranges = append(ranges, strconv.Itoa(i))
	} else {
		ranges = append(ranges, strconv.Itoa(i)+"->"+strconv.Itoa(j))
	}

	return ranges
}
