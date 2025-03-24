package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCanCompleteCircuit(t *testing.T) {
	tests := []struct {
		gas      []int
		cost     []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}, 3},
		{[]int{2, 3, 4}, []int{3, 4, 3}, -1},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("canCompleteCircuit(%v,%v)", test.gas, test.cost), func(t *testing.T) {
			actual := canCompleteCircuit(test.gas, test.cost)
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("expected %d, got %d", test.expected, actual)
			}
		})
	}
}
