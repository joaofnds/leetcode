package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	cur := root
	for cur != nil {
		if cur.Val == val {
			return cur
		}

		if val < cur.Val {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	return nil
}

func searchBSTRecur(root *TreeNode, val int) *TreeNode {
	switch {
	case root == nil:
		return nil
	case val == root.Val:
		return root
	case val < root.Val:
		return searchBSTRecur(root.Left, val)
	case val > root.Val:
		return searchBSTRecur(root.Right, val)
	default:
		return nil
	}
}
