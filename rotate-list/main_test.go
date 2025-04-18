package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRotateRight(t *testing.T) {
	tests := []struct {
		list     []int
		k        int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{4, 5, 1, 2, 3}},
		{[]int{0, 1, 2}, 4, []int{2, 0, 1}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("rotateRight(%v,%d)", test.list, test.k), func(t *testing.T) {
			actual := toSlice(rotateRight(toList(test.list), test.k))
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("expected %v, got %v", test.expected, actual)
			}
		})
	}
}

func toList(nums []int) *ListNode {
	head := &ListNode{Val: nums[0]}
	current := head
	for _, num := range nums[1:] {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return head
}

func toSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}
