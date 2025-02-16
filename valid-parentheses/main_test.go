package main

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		s        string
		expected bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([])", true},
		{"([)]", false},
		{"([}}])", false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("isValid(%q)", test.s), func(t *testing.T) {
			actual := isValid(test.s)
			if actual != test.expected {
				t.Errorf("expected %t, got %t", test.expected, actual)
			}
		})
	}
}
