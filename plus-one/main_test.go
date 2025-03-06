package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		list1    []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{[]int{9}, []int{1, 0}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("plusOne(%v)", test.list1), func(t *testing.T) {
			actual := plusOne(test.list1)
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("expected %v, got %v", test.expected, actual)
			}
		})
	}
}
