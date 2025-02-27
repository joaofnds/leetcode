package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	return &TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBST(nums[:mid]),
		Right: sortedArrayToBST(nums[mid+1:]),
	}
}

func sortedArrayToBSTIter(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	type Task struct {
		Node *TreeNode
		Nums []int
	}

	root := &TreeNode{}
	queue := []Task{{Node: root, Nums: nums}}

	for len(queue) > 0 {
		task := queue[0]
		queue = queue[1:]

		mid := len(task.Nums) / 2
		task.Node.Val = task.Nums[mid]

		if mid > 0 {
			task.Node.Left = &TreeNode{}
			queue = append(queue, Task{Node: task.Node.Left, Nums: task.Nums[:mid]})
		}

		if mid+1 < len(task.Nums) {
			task.Node.Right = &TreeNode{}
			queue = append(queue, Task{Node: task.Node.Right, Nums: task.Nums[mid+1:]})
		}
	}

	return root
}
