package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	return dummy.Next
}

func mergeTwoListsRecur(list1 *ListNode, list2 *ListNode) *ListNode {
	switch {
	case list1 == nil:
		return list2
	case list2 == nil:
		return list1
	case list1.Val < list2.Val:
		list1.Next = mergeTwoListsRecur(list1.Next, list2)
		return list1
	default:
		list2.Next = mergeTwoListsRecur(list1, list2.Next)
		return list2
	}
}
