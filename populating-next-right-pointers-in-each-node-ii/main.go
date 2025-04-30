package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		var prev *Node

		for range len(queue) {
			curr := queue[0]
			queue = queue[1:]

			if prev != nil {
				prev.Next = curr
			}

			prev = curr

			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
	}

	return root
}
