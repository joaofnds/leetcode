package main

import (
	"fmt"
	"testing"
)

func TestWordPattern(t *testing.T) {
	tests := []struct {
		pattern  string
		s        string
		expected bool
	}{
		{"abba", "dog cat cat dog", true},
		{"abba", "dog cat cat fish", false},
		{"aaaa", "dog cat cat dog", false},
		{"abba", "dog dog dog dog", false},
		{"abc", "dog cat dog", false},
		{"aaa", "aa aa aa aa", false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("wordPattern(%q,%q)", test.pattern, test.s), func(t *testing.T) {
			actual := wordPattern(test.pattern, test.s)
			if actual != test.expected {
				t.Errorf("expected %t, got %t", test.expected, actual)
			}
		})
	}
}
