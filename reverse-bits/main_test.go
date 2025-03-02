package main

import (
	"fmt"
	"testing"
)

func TestReverseBits(t *testing.T) {
	tests := []struct {
		num      uint32
		expected uint32
	}{
		{0b00000010100101000001111010011100, 0b00111001011110000010100101000000},
		{0b11111111111111111111111111111101, 0b10111111111111111111111111111111},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("reverseBits(%b)", test.num), func(t *testing.T) {
			actual := reverseBits(test.num)
			if actual != test.expected {
				t.Errorf("expected %b, got %b", test.expected, actual)
			}
		})
	}
}
