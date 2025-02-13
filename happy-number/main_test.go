package main

import (
	"fmt"
	"testing"
)

func TestIsHappy(t *testing.T) {
	tests := []struct {
		n        int
		expected bool
	}{
		{19, true},
		{2, false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("isHappy(%d)", test.n), func(t *testing.T) {
			actual := isHappy(test.n)
			if actual != test.expected {
				t.Errorf("expected %t, got %t", test.expected, actual)
			}
		})
	}
}
