package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGameOfLife(t *testing.T) {
	tests := []struct {
		board    [][]int
		expected [][]int
	}{
		{[][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}, [][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 1, 0}}},
		{[][]int{{1, 1}, {1, 0}}, [][]int{{1, 1}, {1, 1}}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("gameOfLife(%v)", test.board), func(t *testing.T) {
			gameOfLife(test.board)
			if !reflect.DeepEqual(test.board, test.expected) {
				t.Errorf("expected %v, got %v", test.expected, test.board)
			}
		})
	}
}
