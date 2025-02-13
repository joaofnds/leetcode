package main

import (
	"fmt"
	"testing"
)

func TestStrStr(t *testing.T) {
	tests := []struct {
		haystack string
		needle   string
		expected int
	}{
		{"sadbutsad", "sad", 0},
		{"leetcode", "leeto", -1},
		{"mississippi", "issip", 4},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("strStr(%q, %q)", test.haystack, test.needle), func(t *testing.T) {
			actual := strStr(test.haystack, test.needle)
			if actual != test.expected {
				t.Errorf("expected %d, got %d", test.expected, actual)
			}
		})
	}
}
