package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	_, minDiff := inOrder(root, nil, math.MaxInt)
	return minDiff
}

func inOrder(node *TreeNode, prev *TreeNode, minDiff int) (*TreeNode, int) {
	if node == nil {
		return prev, minDiff
	}

	prev, minDiff = inOrder(node.Left, prev, minDiff)

	if prev != nil && node.Val-prev.Val < minDiff {
		minDiff = node.Val - prev.Val
	}

	return inOrder(node.Right, node, minDiff)
}
