package main

import (
	"fmt"
	"testing"
)

func TestPivotIndex(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 7, 3, 6, 5, 6}, 3},
		{[]int{1, 2, 3}, -1},
		{[]int{2, 1, -1}, 0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("pivotIndex(%v)", test.nums), func(t *testing.T) {
			actual := pivotIndex(test.nums)
			if actual != test.expected {
				t.Errorf("expected %d, got %d", test.expected, actual)
			}
		})
	}
}
