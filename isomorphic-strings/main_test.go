package main

import (
	"fmt"
	"testing"
)

func TestIsIsomorphic(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected bool
	}{
		{"egg", "add", true},
		{"foo", "bar", false},
		{"paper", "title", true},
		{"bacd", "baba", false},
		{"aa", "ab", false},
		{"12", "42", true},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("isIsomorphic(%q,%q)", test.s, test.t), func(t *testing.T) {
			actual := isIsomorphic(test.s, test.t)
			if actual != test.expected {
				t.Errorf("expected %t, got %t", test.expected, actual)
			}
		})
	}
}
