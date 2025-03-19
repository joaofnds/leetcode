package main

import (
	"fmt"
	"testing"
)

func TestReverseWords(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"the sky is blue", "blue is sky the"},
		{"  hello world!  ", "world! hello"},
		{"a good   example", "example good a"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("reverseWords(%q)", test.input), func(t *testing.T) {
			actual := reverseWords(test.input)
			if actual != test.expected {
				t.Errorf("expected %q, got %q", test.expected, actual)
			}
		})
	}
}
