package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFlatten(t *testing.T) {
	tests := []struct {
		tree     []int
		expected []int
	}{
		{[]int{1, 2, 5, 3, 4, -1, 6}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{}, []int{}},
		{[]int{0}, []int{0}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("flatten(%v)", test.tree), func(t *testing.T) {
			tree := build(test.tree)
			flatten(tree)
			actual := rightNodesList(tree)

			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("expected %v, got %v", test.expected, actual)
			}
		})
	}
}

func build(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	root := &TreeNode{Val: nums[0]}
	queue := []*TreeNode{root}

	for i := 1; i+1 < len(nums); i += 2 {
		node := queue[0]
		queue = queue[1:]
		left, right := nums[i], nums[i+1]

		if left != -1 {
			node.Left = &TreeNode{Val: left}
			queue = append(queue, node.Left)
		}

		if right != -1 {
			node.Right = &TreeNode{Val: right}
			queue = append(queue, node.Right)
		}
	}

	return root
}

func rightNodesList(root *TreeNode) []int {
	result := []int{}

	for node := root; node != nil; node = node.Right {
		result = append(result, node.Val)
	}

	return result
}
