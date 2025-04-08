package main

type MinStack struct {
	stack []int
	min   []int
}

func Constructor() MinStack {
	return MinStack{stack: []int{}, min: []int{}}
}

func (ms *MinStack) Push(val int) {
	ms.stack = append(ms.stack, val)

	if len(ms.min) == 0 || val <= ms.min[len(ms.min)-1] {
		ms.min = append(ms.min, val)
	}
}

func (ms *MinStack) Pop() {
	top := ms.stack[len(ms.stack)-1]
	ms.stack = ms.stack[:len(ms.stack)-1]

	if len(ms.min) > 0 && top == ms.min[len(ms.min)-1] {
		ms.min = ms.min[:len(ms.min)-1]
	}
}

func (ms *MinStack) Top() int {
	return ms.stack[len(ms.stack)-1]
}

func (ms *MinStack) GetMin() int {
	return ms.min[len(ms.min)-1]
}
