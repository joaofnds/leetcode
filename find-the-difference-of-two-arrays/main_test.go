package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestFindDifference(t *testing.T) {
	tests := []struct {
		nums1    []int
		nums2    []int
		expected [][]int
	}{
		{[]int{1, 2, 3}, []int{2, 4, 6}, [][]int{{1, 3}, {4, 6}}},
		{[]int{1, 2, 3, 3}, []int{1, 1, 2, 2}, [][]int{{3}, {}}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("findDifference(%v,%v)", test.nums1, test.nums2), func(t *testing.T) {
			actual := findDifference(test.nums1, test.nums2)

			if len(actual) != 2 {
				t.Errorf("expected two slices, got %d", len(actual))
				return
			}

			sort.Ints(actual[0])
			sort.Ints(actual[1])
			sort.Ints(test.expected[0])
			sort.Ints(test.expected[1])

			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("expected %v, got %v", test.expected, actual)
			}
		})
	}
}
