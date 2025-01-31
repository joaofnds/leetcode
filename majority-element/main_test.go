package main

import (
	"fmt"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("majorityElement(%v)", test.nums), func(t *testing.T) {
			actual := majorityElement(test.nums)
			if actual != test.expected {
				t.Errorf("majorityElement(%v) got %d, expected %d", test.nums, actual, test.expected)
			}
		})
	}
}
