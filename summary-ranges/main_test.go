package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSummaryRanges(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []string
	}{
		{[]int{0, 1, 2, 4, 5, 7}, []string{"0->2", "4->5", "7"}},
		{[]int{0, 2, 3, 4, 6, 8, 9}, []string{"0", "2->4", "6", "8->9"}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("summaryRanges(%v)", test.nums), func(t *testing.T) {
			actual := summaryRanges(test.nums)
			if !reflect.DeepEqual(test.expected, actual) {
				t.Errorf("expected %v, got %v", test.expected, actual)
			}
		})
	}
}
