package main

import (
	"fmt"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"ab", "a", false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("isAnagram(%q,%q)", test.s, test.t), func(t *testing.T) {
			actual := isAnagram(test.s, test.t)
			if actual != test.expected {
				t.Errorf("expected %t, got %t", actual, test.expected)
			}
		})
	}
}
