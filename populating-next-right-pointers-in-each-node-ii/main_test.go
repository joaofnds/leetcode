package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConnect(t *testing.T) {
	tests := []struct {
		tree     []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5, -1, 7}, []int{1, 2, 3, 4, 5, 7}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("connect(%v)", test.tree), func(t *testing.T) {
			actual := byLevelUsingNext(connect(build(test.tree)))
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("expected %v, got %v", test.expected, actual)
			}
		})
	}
}

func build(nums []int) *Node {
	if len(nums) == 0 {
		return nil
	}

	if len(nums)%2 == 0 {
		nums = append(nums, -1)
	}

	root := &Node{Val: nums[0]}
	queue := []*Node{root}

	for i := 1; i+1 < len(nums); i += 2 {
		node := queue[0]
		queue = queue[1:]
		left, right := nums[i], nums[i+1]

		if left != -1 {
			node.Left = &Node{Val: left}
			queue = append(queue, node.Left)
		}

		if right != -1 {
			node.Right = &Node{Val: right}
			queue = append(queue, node.Right)
		}
	}

	return root
}

func byLevelUsingNext(root *Node) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	firstOfLevel := root

	for firstOfLevel != nil {
		for current := firstOfLevel; current != nil; current = current.Next {
			result = append(result, current.Val)
		}
		firstOfLevel = firstOfLevel.Left
	}

	return result
}
