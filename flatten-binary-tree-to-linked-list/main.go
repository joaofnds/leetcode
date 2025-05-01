package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	for curr := root; curr != nil; curr = curr.Right {
		if curr.Left == nil {
			continue
		}

		rightmost := curr.Left
		for rightmost.Right != nil {
			rightmost = rightmost.Right
		}

		rightmost.Right = curr.Right
		curr.Right = curr.Left
		curr.Left = nil
	}
}
