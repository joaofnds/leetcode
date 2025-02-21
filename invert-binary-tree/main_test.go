package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInvertTree(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{[]int{2, 1, 3}, []int{2, 3, 1}},
		{[]int{4, 2, 7, 1, 3, 6, 9}, []int{4, 7, 2, 9, 6, 3, 1}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("invertTree(%v)", test.nums), func(t *testing.T) {
			actual := treeToList(invertTree(listToTree(test.nums)))
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("expected %v, got %v", test.expected, actual)
			}
		})
	}
}

func listToTree(nums []int) *TreeNode {
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

func treeToList(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			result = append(result, -1)
		} else {
			result = append(result, node.Val)
			queue = append(queue, node.Left, node.Right)
		}
	}

	for len(result) > 0 && result[len(result)-1] == -1 {
		result = result[:len(result)-1]
	}

	return result
}
