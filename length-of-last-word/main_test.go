package main

import (
	"fmt"
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"Hello World", 5},
		{"   fly me   to   the moon  ", 4},
		{"luffy is still joyboy", 6},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("lengthOfLastWord(%q)", test.input), func(t *testing.T) {
			actual := lengthOfLastWord(test.input)
			if actual != test.expected {
				t.Errorf("lengthOfLastWord(%q) got %d, expected %d", test.input, actual, test.expected)
			}
		})
	}
}
