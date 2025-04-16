package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		list     []int
		n        int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5}},
		{[]int{1}, 1, []int{}},
		{[]int{1, 2}, 1, []int{1}},
		{[]int{4, 6, 8, 8, 2, 1, 3, 6, 6, 0, 6, 2, 0, 3, 8, 7, 6, 0, 3, 5, 1, 9, 5, 7, 4}, 21, []int{4, 6, 8, 8, 1, 3, 6, 6, 0, 6, 2, 0, 3, 8, 7, 6, 0, 3, 5, 1, 9, 5, 7, 4}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("removeNthFromEnd(%v,%d)", test.list, test.n), func(t *testing.T) {
			actual := removeNthFromEnd(toList(test.list), test.n)
			actualAsSlice := toSlice(actual)
			if !reflect.DeepEqual(actualAsSlice, test.expected) {
				t.Errorf("expected %v, got %v", test.expected, actualAsSlice)
			}
		})
	}
}

func toList(nums []int) *ListNode {
	var head *ListNode
	var current *ListNode
	for _, n := range nums {
		if head == nil {
			head = &ListNode{Val: n}
			current = head
		} else {
			current.Next = &ListNode{Val: n}
			current = current.Next
		}
	}
	return head
}

func toSlice(head *ListNode) []int {
	nums := []int{}
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	return nums
}
