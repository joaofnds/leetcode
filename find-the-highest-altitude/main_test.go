package main

import (
	"fmt"
	"testing"
)

func TestLargestAltitude(t *testing.T) {
	tests := []struct {
		gain     []int
		expected int
	}{
		{[]int{-5, 1, 5, 0, -7}, 1},
		{[]int{-4, -3, -2, -1, 4, 3, 2}, 0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("largestAltitude(%v)", test.gain), func(t *testing.T) {
			actual := largestAltitude(test.gain)
			if actual != test.expected {
				t.Errorf("expected %d, got %d", test.expected, actual)
			}
		})
	}
}
