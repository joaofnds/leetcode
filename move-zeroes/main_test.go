package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{0}, []int{0}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("moveZeroes(%v)", test.nums), func(t *testing.T) {
			moveZeroes(test.nums)
			if !reflect.DeepEqual(test.nums, test.expected) {
				t.Errorf("expected %v, got %v", test.expected, test.nums)
			}
		})
	}
}
