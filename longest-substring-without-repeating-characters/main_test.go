package main

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("lengthOfLongestSubstring(%q)", test.input), func(t *testing.T) {
			actual := lengthOfLongestSubstring(test.input)
			if actual != test.expected {
				t.Errorf("expected %d, got %d", test.expected, actual)
			}
		})
	}
}
