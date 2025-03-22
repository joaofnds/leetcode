package main

func maxArea(heights []int) int {
	var maxArea int

	left, right := 0, len(heights)-1
	for left < right {
		widght, height := right-left, min(heights[left], heights[right])
		maxArea = max(maxArea, widght*height)
		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}
