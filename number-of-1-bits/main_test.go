package main

import (
	"fmt"
	"testing"
)

func TestHammingWeight(t *testing.T) {
	tests := []struct {
		num      int
		expected int
	}{
		{11, 3},
		{128, 1},
		{2147483645, 30},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("hammingWeight(%b)", test.num), func(t *testing.T) {
			actual := hammingWeight(test.num)
			if actual != test.expected {
				t.Errorf("expected %b, got %b", test.expected, actual)
			}
		})
	}
}
