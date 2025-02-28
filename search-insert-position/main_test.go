package main

import (
	"fmt"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected int
	}{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("searchInsert(%v,%d)", test.nums, test.target), func(t *testing.T) {
			actual := searchInsert(test.nums, test.target)
			if actual != test.expected {
				t.Errorf("expected %d, got %d", test.expected, actual)
			}
		})
	}
}
