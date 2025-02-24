package main

import (
	"fmt"
	"testing"
)

func TestCountNodes(t *testing.T) {
	tests := []struct {
		tree     []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, 6},
		{[]int{}, 0},
		{[]int{1}, 1},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("countNodes(%v)", test.tree), func(t *testing.T) {
			actual := countNodes(build(test.tree))
			if actual != test.expected {
				t.Errorf("expected %d, got %d", test.expected, actual)
			}
		})
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("countNodesRecur(%v)", test.tree), func(t *testing.T) {
			actual := countNodesRecur(build(test.tree))
			if actual != test.expected {
				t.Errorf("expected %d, got %d", test.expected, actual)
			}
		})
	}
}

func BenchmarkCountNodes(b *testing.B) {
	for _, size := range []int{10, 100, 1_000, 10_000, 100_000} {
		tree := buildTree(size)

		b.Run(fmt.Sprintf("iterative_%d_nodes", size), func(b *testing.B) {
			for b.Loop() {
				countNodes(tree)
			}
		})

		b.Run(fmt.Sprintf("recursive_%d_nodes", size), func(b *testing.B) {
			for b.Loop() {
				countNodesRecur(tree)
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
