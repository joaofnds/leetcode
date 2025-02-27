package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSortedArrayToBST(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{-10, -3, 0, 5, 9}, []int{0, -3, 9, -10, 5}},
		{[]int{1, 3}, []int{3, 1}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("sortedArrayToBST(%v)", test.input), func(t *testing.T) {
			actual := toSlice(sortedArrayToBST(test.input))
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("expected %v, got %v", test.expected, actual)
			}
		})
	}
}

func toSlice(root *TreeNode) []int {
	result := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node != nil {
			result = append(result, node.Val)
			queue = append(queue, node.Left, node.Right)
		}
	}

	return result
}
