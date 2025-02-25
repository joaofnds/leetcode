package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}

	var result []float64
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSum := 0.0
		levelSize := len(queue)

		for range levelSize {
			node := queue[0]
			queue = queue[1:]
			levelSum += float64(node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, levelSum/float64(levelSize))
	}

	return result
}
