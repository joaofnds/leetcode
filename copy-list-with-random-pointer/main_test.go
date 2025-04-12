package main

import (
	"fmt"
	"testing"
)

func TestCopyRandomList(t *testing.T) {
	inputs := [][][2]int{
		{{7, -1}, {13, 0}, {11, 4}, {10, 2}, {1, 0}},
		{{1, 1}, {2, 1}},
		{{3, -1}, {3, 0}, {3, -1}},
	}

	for _, listAsSlice := range inputs {
		t.Run(fmt.Sprintf("copyRandomList(%v)", listAsSlice), func(t *testing.T) {
			original := toList(listAsSlice)
			clone := copyRandomList(original)

			for original != nil {
				if clone == nil {
					t.Fatalf("expected non-nil node, got nil")
				}

				if original.Val != clone.Val {
					t.Errorf("expected value %d, got %d", original.Val, clone.Val)
				}

				if original.Next != nil {
					if clone.Next == nil {
						t.Errorf("expected non-nil next node, got nil")
					} else if original.Next.Val != clone.Next.Val {
						t.Errorf("expected next value %d, got %d", original.Next.Val, clone.Next.Val)
					}
				} else {
					if clone.Next != nil {
						t.Errorf("expected nil next node, got %d", clone.Next.Val)
					}
				}

				if original.Random != nil {
					if clone.Random == nil {
						t.Errorf("expected non-nil random node, got nil")
					} else if original.Random.Val != clone.Random.Val {
						t.Errorf("expected random value %d, got %d", original.Random.Val, clone.Random.Val)
					}
				} else {
					if clone.Random != nil {
						t.Errorf("expected nil random node, got %d", clone.Random.Val)
					}
				}

				original = original.Next
				clone = clone.Next
			}
		})
	}
}

func toList(nums [][2]int) *Node {
	head := &Node{}
	current := head
	nodes := make(map[int]*Node)

	for i, pair := range nums {
		node := &Node{Val: pair[0]}
		nodes[i] = node
		current.Next = node
		current = node
	}

	for i, pair := range nums {
		if pair[1] != -1 {
			nodes[i].Random = nodes[pair[1]]
		}
	}

	return head.Next
}
