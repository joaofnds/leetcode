package main

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		n        int
		expected bool
	}{
		{n: 121, expected: true},
		{n: -121, expected: false},
		{n: 10, expected: false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("isPalindrome(%d)", test.n), func(t *testing.T) {
			actual := isPalindrome(test.n)
			if actual != test.expected {
				t.Errorf("expected %t, got %t", test.expected, actual)
			}
		})
	}
}
