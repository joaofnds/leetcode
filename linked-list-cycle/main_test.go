package main

import (
	"fmt"
	"testing"
)

func TestHasCycle(t *testing.T) {
	tests := []struct {
		list     []int
		pos      int
		expected bool
	}{
		{[]int{3, 2, 0, -4}, 1, true},
		{[]int{1, 2}, 0, true},
		{[]int{1}, -1, false},
		{[]int{-21, 10, 17, 8, 4, 26, 5, 35, 33, -7, -16, 27, -12, 6, 29, -12, 5, 9, 20, 14, 14, 2, 13, -24, 21, 23, -21, 5}, -1, false},
		{[]int{3, 2, 0, -4}, 1, true},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("hasCycle(%v)", test.list), func(t *testing.T) {
			actual := hasCycle(build(test.list, test.pos))
			if actual != test.expected {
				t.Errorf("expected %t, got %t", test.expected, actual)
			}
		})
	}
}

func build(nums []int, pos int) *ListNode {
	var head, cycle *ListNode
	var prev *ListNode

	for i, n := range nums {
		node := &ListNode{Val: n}
		if i == pos {
			cycle = node
		}
		if prev != nil {
			prev.Next = node
		} else {
			head = node
		}
		prev = node
	}

	if prev != nil {
		prev.Next = cycle
	}

	return head
}
