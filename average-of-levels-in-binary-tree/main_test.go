package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCountNodes(t *testing.T) {
	tests := []struct {
		tree     []int
		expected []float64
	}{
		{[]int{3, 9, 20, -1, -1, 15, 7}, []float64{3, 14.5, 11}},
		{[]int{3, 9, 20, 15, 7}, []float64{3, 14.5, 11}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("averageOfLevels(%v)", test.tree), func(t *testing.T) {
			actual := averageOfLevels(build(test.tree))
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
