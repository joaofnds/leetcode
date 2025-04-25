package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCountNodes(t *testing.T) {
	tests := []struct {
		preorder []int
		inorder  []int
		expected []any
	}{
		{[]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}, []any{3, 9, 20, nil, nil, 15, 7}},
		{[]int{-1}, []int{-1}, []any{-1}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("buildTree(%v,%v)", test.preorder, test.inorder), func(t *testing.T) {
			root := buildTree(test.preorder, test.inorder)
			result := levelOrder(root)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

func levelOrder(root *TreeNode) []any {
	if root == nil {
		return nil
	}

	var result []any
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node != nil {
			result = append(result, node.Val)
			queue = append(queue, node.Left, node.Right)
		} else {
			result = append(result, nil)
		}
	}

	for len(result) > 0 && result[len(result)-1] == nil {
		result = result[:len(result)-1]
	}

	return result
}
