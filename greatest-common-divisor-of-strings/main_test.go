package main

import (
	"fmt"
	"testing"
)

func TestGCDOfStrings(t *testing.T) {
	tests := []struct {
		a        string
		b        string
		expected string
	}{
		{"ABCABC", "ABC", "ABC"},
		{"ABABAB", "ABAB", "AB"},
		{"LEET", "CODE", ""},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("gcdOfStrings(%q,%q)", test.a, test.b), func(t *testing.T) {
			actual := gcdOfStrings(test.a, test.b)
			if actual != test.expected {
				t.Errorf("expected %q, got %q", test.expected, actual)
			}
		})
	}
}
