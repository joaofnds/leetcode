package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	tests := []struct {
		target   int
		nums     []int
		expected int
	}{
		{7, []int{2, 3, 1, 2, 4, 3}, 2},
		{4, []int{1, 4, 4}, 1},
		{11, []int{1, 1, 1, 1, 1, 1, 1, 1}, 0},
		{11, []int{1, 2, 3, 4, 5}, 3},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("minSubArrayLen(%d,%v)", test.target, test.nums), func(t *testing.T) {
			actual := minSubArrayLen(test.target, test.nums)
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("expected %d, got %d", test.expected, actual)
			}
		})
	}
}
