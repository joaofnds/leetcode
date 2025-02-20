package main

import (
	"fmt"
	"testing"
)

func TestIsSameTree(t *testing.T) {
	tests := []struct {
		p        []int
		q        []int
		expected bool
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}, true},
		{[]int{1, 2}, []int{1, -1, 2}, false},
		{[]int{1, 2, 1}, []int{1, 1, 2}, false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("isSameTree(%v,%v)", test.p, test.q), func(t *testing.T) {
			actual := isSameTree(build(test.p), build(test.q))
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
