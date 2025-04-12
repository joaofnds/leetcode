package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	// create a copy of each node and insert it next to the original node
	current := head
	for current != nil {
		newNode := &Node{Val: current.Val}
		newNode.Next = current.Next
		current.Next = newNode
		current = newNode.Next
	}

	// set the random pointers for the copied nodes
	current = head
	for current != nil {
		if current.Random != nil {
			current.Next.Random = current.Random.Next
		}
		current = current.Next.Next
	}

	// separate the two lists
	current = head
	newHead := head.Next
	for current != nil {
		newNode := current.Next
		current.Next = newNode.Next
		if newNode.Next != nil {
			newNode.Next = newNode.Next.Next
		}
		current = current.Next
	}

	return newHead
}
