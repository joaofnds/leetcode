package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPartition(t *testing.T) {
	tests := []struct {
		list     []int
		x        int
		expected []int
	}{
		{[]int{1, 4, 3, 2, 5, 2}, 3, []int{1, 2, 2, 4, 3, 5}},
		{[]int{2, 1}, 2, []int{1, 2}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("partition(%v,%d)", test.list, test.x), func(t *testing.T) {
			actual := toSlice(partition(toList(test.list), test.x))
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
