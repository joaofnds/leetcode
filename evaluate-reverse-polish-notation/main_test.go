package main

import (
	"fmt"
	"testing"
)

func TestEvalRPN(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{[]string{"2", "1", "+", "3", "*"}, 9},
		{[]string{"4", "13", "5", "/", "+"}, 6},
		{[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("evalRPN(%v)", test.input), func(t *testing.T) {
			actual := evalRPN(test.input)
			if actual != test.expected {
				t.Errorf("expected %d, got %d", test.expected, actual)
			}
		})
	}
}
