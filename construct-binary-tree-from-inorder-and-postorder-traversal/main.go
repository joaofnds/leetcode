package main

import "slices"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	i := slices.Index(inorder, postorder[len(postorder)-1])

	return &TreeNode{
		Val:   postorder[len(postorder)-1],
		Left:  buildTree(inorder[:i], postorder[:i]),
		Right: buildTree(inorder[i+1:], postorder[i:len(postorder)-1]),
	}
}
