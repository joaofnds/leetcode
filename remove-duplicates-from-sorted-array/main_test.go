package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{[]int{1, 1, 2}, []int{1, 2}},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("removeDuplicates(%v)", test.nums), func(t *testing.T) {
			k := removeDuplicates(test.nums)
			if !reflect.DeepEqual(test.nums[:k], test.expected) {
				t.Errorf("expected %v, got %v", test.expected, test.nums[:k])
			}
		})
	}
}
