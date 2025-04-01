package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetZeroes(t *testing.T) {
	tests := []struct {
		matrix   [][]int
		expected [][]int
	}{
		{[][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}},
		{[][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}, [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("setZeroes(%v)", test.matrix), func(t *testing.T) {
			setZeroes(test.matrix)
			if !reflect.DeepEqual(test.matrix, test.expected) {
				t.Errorf("expected %v, got %v", test.expected, test.matrix)
			}
		})
	}
}
