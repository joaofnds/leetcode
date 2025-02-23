package main

import (
	"fmt"
	"testing"
)

func TestHasPathSum(t *testing.T) {
	tests := []struct {
		tree      []int
		targetSum int
		expected  bool
	}{
		{[]int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, -1, 1}, 22, true},
		{[]int{1, 2, 3}, 5, false},
		{[]int{}, 0, false},
		{[]int{1, 2}, 1, false},
		{[]int{10, 2, 11, 0}, 12, true},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("hasPathSum(%v, %d)", test.tree, test.targetSum), func(t *testing.T) {
			actual := hasPathSum(build(test.tree), test.targetSum)
			if actual != test.expected {
				t.Errorf("expected %t, got %t", test.expected, actual)
			}
		})
	}
}

func build(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	if len(nums)%2 == 0 {
		nums = append(nums, -1)
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
