package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	// find the length of the list
	length := 1
	current := head
	for current.Next != nil {
		current = current.Next
		length++
	}

	// connect the tail to the head to form a circle
	current.Next = head

	// find the new tail
	newTail := length - k%length - 1
	current = head
	for i := 0; i < newTail; i++ {
		current = current.Next
	}

	// set the new head and break the circle
	newHead := current.Next
	current.Next = nil

	return newHead
}
