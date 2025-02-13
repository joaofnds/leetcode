package main

import (
	"fmt"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"III", 3},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("romanToInt(%q)", test.input), func(t *testing.T) {
			actual := romanToInt(test.input)
			if actual != test.expected {
				t.Errorf("expected %d, got %d", test.expected, actual)
			}
		})
	}
}
