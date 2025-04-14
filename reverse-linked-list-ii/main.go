package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}

	dummy := &ListNode{Next: head}

	pre := dummy
	for i := 1; i < left; i++ {
		pre = pre.Next
	}

	reverseStart := pre.Next

	reverseEnd := reverseStart
	for i := left; i < right; i++ {
		reverseEnd = reverseEnd.Next
	}

	next := reverseEnd.Next
	reverseEnd.Next = nil

	pre.Next = reverse(reverseStart)
	reverseStart.Next = next

	return dummy.Next
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	current := head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}
