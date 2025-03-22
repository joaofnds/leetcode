package main

import (
	"fmt"
	"testing"
)

func TestMaxArea(t *testing.T) {
	tests := []struct {
		height   []int
		expected int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
    {[]int{1, 1}, 1},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("maxArea(%v)", test.height), func(t *testing.T) {
			actual := maxArea(test.height)
			if actual != test.expected {
				t.Errorf("expected %d, got %d", test.expected, actual)
			}
		})
	}
}
