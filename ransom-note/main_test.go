package main

import (
	"fmt"
	"testing"
)

func TestCanConstruct(t *testing.T) {
	tests := []struct {
		ransomNote string
		magazine   string
		expected   bool
	}{
		{"a", "b", false},
		{"aa", "ab", false},
		{"aa", "aab", true},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("canConstruct(%q,%q)", test.ransomNote, test.magazine), func(t *testing.T) {
			actual := canConstruct(test.ransomNote, test.magazine)
			if actual != test.expected {
				t.Errorf("expected %t, got %t", test.expected, actual)
			}
		})
	}
}
