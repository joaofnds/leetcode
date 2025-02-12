package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("twoSum(%v,%d)", test.nums, test.target), func(t *testing.T) {
			actual := twoSum(test.nums, test.target)
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("expected %v, got %v", actual, test.expected)
			}
		})
	}
}
