package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		nums     []int
		val      int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},

		{[]int{1, 2, 3}, 0, []int{1, 2, 3}},
		{[]int{1, 2, 3}, 1, []int{3, 1, 2}},
		{[]int{1, 2, 3}, 2, []int{2, 3, 1}},
		{[]int{1, 2, 3}, 3, []int{1, 2, 3}},

		{[]int{1, 2, 3, 4, 5, 6, 7}, 0, []int{1, 2, 3, 4, 5, 6, 7}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 7, []int{1, 2, 3, 4, 5, 6, 7}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 14, []int{1, 2, 3, 4, 5, 6, 7}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("rotate(%v, %d)", test.nums, test.val), func(t *testing.T) {
			rotate(test.nums, test.val)
			if !reflect.DeepEqual(test.nums, test.expected) {
				t.Errorf("expected %v, got %v", test.expected, test.nums)
			}
		})
	}
}
