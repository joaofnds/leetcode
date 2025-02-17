package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	var prev *ListNode
	start := head

	for head != nil {
		next := head.Next
		if next == start {
			return true
		}
		head.Next = prev
		prev = head
		head = next
	}

	return false
}
