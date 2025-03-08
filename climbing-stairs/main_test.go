package main

import (
	"fmt"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{2, 2},
		{3, 3},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("climbStairs(%d)", test.input), func(t *testing.T) {
			actual := climbStairs(test.input)
			if actual != test.expected {
				t.Errorf("expected %d, got %d", test.expected, actual)
			}
		})
	}
}
