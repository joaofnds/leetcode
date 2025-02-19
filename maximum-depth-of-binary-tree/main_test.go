package main

import (
	"fmt"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		tree     []int
		expected int
	}{
		{[]int{3, 9, 20, -1, -1, 15, 7}, 3},
		{[]int{1, -1, 2}, 2},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("maxDepth(%v)", test.tree), func(t *testing.T) {
			actual := maxDepth(build(test.tree))
			if actual != test.expected {
				t.Errorf("expected %d, got %d", test.expected, actual)
			}
		})
	}
}

// given a list like [3,9,20,null,null,15,7], build a tree like:
//
//	  3
//	 / \
//	9  20
//	  /  \
//	 15  7
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
