package main

import (
	"fmt"
	"testing"
)

func TestFindMaxAverage(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected float64
	}{
		{[]int{1, 12, -5, -6, 50, 3}, 4, 12.75},
		{[]int{5}, 1, 5.0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("findMaxAverage(%v,%d)", test.nums, test.k), func(t *testing.T) {
			actual := findMaxAverage(test.nums, test.k)
			if actual != test.expected {
				t.Errorf("expected %f, got %f", test.expected, actual)
			}
		})
	}
}
