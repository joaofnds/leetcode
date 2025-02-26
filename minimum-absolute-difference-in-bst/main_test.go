package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetMinimumDifference(t *testing.T) {
	tests := []struct {
		tree     []int
		expected int
	}{
		{[]int{4, 2, 6, 1, 3}, 1},
		{[]int{1, 0, 48, -1, -1, 12, 49}, 1},
		{[]int{5, 4, 7}, 1},
		{[]int{236, 104, 701, -1, 227, -1, 911}, 9},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("getMinimumDifference(%v)", test.tree), func(t *testing.T) {
			actual := getMinimumDifference(build(test.tree))
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
