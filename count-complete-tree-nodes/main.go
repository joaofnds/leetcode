package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	queue := []*TreeNode{root}
	var count int

	for len(queue) > 0 {
		root = queue[0]
		queue = queue[1:]

		if root != nil {
			count++
			queue = append(queue, root.Left, root.Right)
		}
	}

	return count
}

func countNodesRecur(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + countNodesRecur(root.Left) + countNodesRecur(root.Right)
}
