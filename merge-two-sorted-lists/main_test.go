package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		list1    []int
		list2    []int
		expected []int
	}{
		{[]int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
		{[]int{}, []int{}, []int{}},
		{[]int{}, []int{0}, []int{0}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("mergeTwoLists(%v,%v)", test.list1, test.list2), func(t *testing.T) {
			actual := mergeTwoLists(toList(test.list1), toList(test.list2))
			actualAsSlice := toSlice(actual)
			if !reflect.DeepEqual(actualAsSlice, test.expected) {
				t.Errorf("expected %v, got %v", test.expected, actualAsSlice)
			}
		})
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("mergeTwoListsRecur(%v,%v)", test.list1, test.list2), func(t *testing.T) {
			actual := mergeTwoListsRecur(toList(test.list1), toList(test.list2))
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
