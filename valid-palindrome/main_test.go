package main

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{" ", true},
		{"0P", false},
		{"Doc, note: I dissent. A fast never prevents a fatness. I diet on cod", true},
		{"T. Eliot, top bard, notes putrid tang emanating, is sad; I'd assign it a name: gnat dirt upset on drab pot toilet.", true},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("isPalindrome(%q)", test.input), func(t *testing.T) {
			actual := isPalindrome(test.input)
			if actual != test.expected {
				t.Errorf("isPalindrome(%q) got %t, expected %t", test.input, actual, test.expected)
			}
		})
	}
}
