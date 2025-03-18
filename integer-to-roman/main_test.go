package main

import (
	"fmt"
	"testing"
)

func TestIntToRoman(t *testing.T) {
	tests := []struct {
		num      int
		expected string
	}{
		{3749, "MMMDCCXLIX"},
		{58, "LVIII"},
		{1994, "MCMXCIV"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("intToRoman(%d)", test.num), func(t *testing.T) {
			actual := intToRoman(test.num)
			if actual != test.expected {
				t.Errorf("expected %q, got %q", test.expected, actual)
			}
		})
	}
}
