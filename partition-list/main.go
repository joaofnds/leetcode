package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	dummy1 := &ListNode{}
	dummy2 := &ListNode{}
	p1, p2 := dummy1, dummy2

	for head != nil {
		if head.Val < x {
			p1.Next = head
			p1 = p1.Next
		} else {
			p2.Next = head
			p2 = p2.Next
		}
		head = head.Next
	}

	p2.Next = nil
	p1.Next = dummy2.Next

	return dummy1.Next
}
