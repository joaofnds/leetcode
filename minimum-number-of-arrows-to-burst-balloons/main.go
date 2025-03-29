package main

import (
	"math"
	"sort"
)

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	count, end := 0, math.MinInt
	for _, point := range points {
		if point[0] > end {
			count++
			end = point[1]
		}
	}
	return count
}
