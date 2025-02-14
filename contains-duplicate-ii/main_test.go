package main

import (
	"fmt"
	"testing"
)

func TestContainsNearbyDuplicate(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected bool
	}{
		{[]int{1, 2, 3, 1}, 3, true},
		{[]int{1, 0, 1, 1}, 1, true},
		{[]int{1, 2, 3, 1, 2, 3}, 2, false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("containsNearbyDuplicate(%v,%d)", test.nums, test.k), func(t *testing.T) {
			actual := containsNearbyDuplicate(test.nums, test.k)
			if actual != test.expected {
				t.Errorf("expected %t, got %t", test.expected, actual)
			}
		})
	}
}
