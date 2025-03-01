package main

import (
	"fmt"
	"testing"
)

func TestAddBinary(t *testing.T) {
	tests := []struct {
		a        string
		b        string
		expected string
	}{
		{"0", "0", "0"},
		{"11", "1", "100"},
		{"1010", "1011", "10101"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("addBinary(%q,%q)", test.a, test.b), func(t *testing.T) {
			actual := addBinary(test.a, test.b)
			if actual != test.expected {
				t.Errorf("expected %q, got %q", test.expected, actual)
			}
		})
	}
}
