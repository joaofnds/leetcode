package main

import (
	"fmt"
	"math"
	"testing"
)

func TestMySqrt(t *testing.T) {
	for x := range 100 {
		expected := int(math.Sqrt(float64(x)))

		t.Run(fmt.Sprintf("mySqrt(%d)", x), func(t *testing.T) {
			actual := mySqrt(x)
			if actual != expected {
				t.Errorf("expected %d, got %d", expected, actual)
			}
		})
	}
}
