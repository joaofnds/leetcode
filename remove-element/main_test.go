package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		nums     []int
		val      int
		expected []int
	}{
		{[]int{3, 2, 2, 3}, 3, []int{2, 2}},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, []int{0, 1, 3, 0, 4}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("removeElement(%v, %d)", test.nums, test.val), func(t *testing.T) {
			k := removeElement(test.nums, test.val)
			if !reflect.DeepEqual(test.nums[:k], test.expected) {
				t.Errorf("expected %v, got %v", test.expected, test.nums[:k])
			}
		})
	}
}
