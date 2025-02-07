package main

import (
	"fmt"
	"testing"
)

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected bool
	}{
		{"abc", "ahbgdc", true},
		{"axc", "ahbgdc", false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("isSubsequence(%q, %q)", test.s, test.t), func(t *testing.T) {
			actual := isSubsequence(test.s, test.t)
			if actual != test.expected {
				t.Errorf("isSubsequence(%q, %q) got %t, expected %t", test.s, test.t, actual, test.expected)
			}
		})
	}
}
