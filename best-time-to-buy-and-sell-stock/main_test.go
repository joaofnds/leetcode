package main

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("maxProfit(%v)", test.nums), func(t *testing.T) {
			actual := maxProfit(test.nums)
			if actual != test.expected {
				t.Errorf("maxProfit(%v) got %d, expected %d", test.nums, actual, test.expected)
			}
		})
	}
}
