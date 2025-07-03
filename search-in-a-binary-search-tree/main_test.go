package main

import (
	"fmt"
	"testing"
)

func TestSearchBST(t *testing.T) {
	tests := []struct {
		tree     []int
		val      int
		expected *TreeNode
	}{
		{[]int{4, 2, 7, 1, 3}, 2, &TreeNode{Val: 2}},
		{[]int{4, 2, 7, 1, 3}, 5, nil},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("searchBST(%v)", test.tree), func(t *testing.T) {
			actual := searchBST(build(test.tree), test.val)
			if !isSameNodeValue(actual, test.expected) {
				t.Errorf("expected %v, got %v", test.expected, actual)
			}
		})
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("searchBSTRecur(%v)", test.tree), func(t *testing.T) {
			actual := searchBSTRecur(build(test.tree), test.val)
			if !isSameNodeValue(actual, test.expected) {
				t.Errorf("expected %v, got %v", test.expected, actual)
			}
		})
	}
}

func BenchmarkSearchBST(b *testing.B) {
	for _, size := range []int{10, 100, 1_000, 10_000, 100_000} {
		tree := buildTree(size)
		searchVal := size/2 + 1
		println(searchVal)

		b.Run(fmt.Sprintf("iterative_%d_nodes", size), func(b *testing.B) {
			for b.Loop() {
				searchBST(tree, searchVal)
			}
		})

		b.Run(fmt.Sprintf("recursive_%d_nodes", size), func(b *testing.B) {
			for b.Loop() {
				searchBSTRecur(tree, searchVal)
			}
		})
	}

	for _, size := range []int{10, 100, 1_000, 10_000, 100_000} {
		tree := buildTree(size)
		searchVal := size/2 + 1

		b.Run(fmt.Sprintf("iterative_%d_nodes", size), func(b *testing.B) {
			for b.Loop() {
				searchBST(tree, searchVal)
			}
		})

		b.Run(fmt.Sprintf("recursive_%d_nodes", size), func(b *testing.B) {
			for b.Loop() {
				searchBSTRecur(tree, searchVal)
			}
		})
	}
}

func isSameNodeValue(a *TreeNode, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	return a.Val == b.Val
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

func buildTree(size int) *TreeNode {
	if size == 0 {
		return nil
	}

	root := &TreeNode{Val: 1}
	queue := []*TreeNode{root}

	for i := 2; i <= size; i++ {
		node := queue[0]
		queue = queue[1:]

		node.Left = &TreeNode{Val: i}
		queue = append(queue, node.Left)

		if i < size {
			node.Right = &TreeNode{Val: i + 1}
			queue = append(queue, node.Right)
			i++
		}
	}

	return root
}
