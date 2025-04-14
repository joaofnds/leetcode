package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReverseBetween(t *testing.T) {
	tests := []struct {
		list     []int
		left     int
		right    int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, 4, []int{1, 4, 3, 2, 5}},
		{[]int{5}, 1, 1, []int{5}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("reverseBetween(%v,%d,%d)", test.list, test.left, test.right), func(t *testing.T) {
			actual := reverseBetween(toList(test.list), test.left, test.right)
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
