package main

import (
	"fmt"
	"testing"
)

func TestConvert(t *testing.T) {
	tests := []struct {
		str      string
		rows     int
		expected string
	}{
		{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
		{"A", 1, "A"},
		{"AB", 1, "AB"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("convert(%q,%d)", test.str, test.rows), func(t *testing.T) {
			actual := convert(test.str, test.rows)
			if actual != test.expected {
				t.Errorf("expected %q, got %q", test.expected, actual)
			}
		})
	}
}
