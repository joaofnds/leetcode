package main

import (
	"fmt"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		inputs   []string
		expected string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("longestCommonPrefix(%v)", test.inputs), func(t *testing.T) {
			actual := longestCommonPrefix(test.inputs)
			if actual != test.expected {
				t.Errorf("expected %q, got %q", test.expected, actual)
			}
		})
	}
}
