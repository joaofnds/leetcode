package main

import (
	"fmt"
	"testing"
)

func TestJump(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{2, 3, 1, 1, 4}, 2},
		{[]int{2, 3, 0, 1, 4}, 2},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("jump(%v)", test.nums), func(t *testing.T) {
			actual := jump(test.nums)
			if actual != test.expected {
				t.Errorf("expected %d, got %d", test.expected, actual)
			}
		})
	}
}
