package main

import (
	"fmt"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
		{[]int{1}, 1},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("singleNumber(%v)", test.nums), func(t *testing.T) {
			actual := singleNumber(test.nums)
			if actual != test.expected {
				t.Errorf("expected %d, got %d", test.expected, actual)
			}
		})
	}
}
