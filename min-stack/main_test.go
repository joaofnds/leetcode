package main

import (
	"testing"
)

func TestMinStack(t *testing.T) {
	t.Run("Push, Pop, Top, GetMin operations", func(t *testing.T) {
		minStack := Constructor()

		// Push elements
		minStack.Push(-2)
		minStack.Push(0)
		minStack.Push(-3)

		// Test GetMin
		if got := minStack.GetMin(); got != -3 {
			t.Errorf("GetMin() = %v, want %v", got, -3)
		}

		// Test Pop
		minStack.Pop()

		// Test Top
		if got := minStack.Top(); got != 0 {
			t.Errorf("Top() = %v, want %v", got, 0)
		}

		// Test GetMin after pop
		if got := minStack.GetMin(); got != -2 {
			t.Errorf("GetMin() = %v, want %v", got, -2)
		}
	})

	t.Run("Empty stack operations", func(t *testing.T) {
		minStack := Constructor()

		// Push and then pop everything
		minStack.Push(5)
		minStack.Pop()

		// Push again after empty
		minStack.Push(10)
		if got := minStack.Top(); got != 10 {
			t.Errorf("Top() = %v, want %v", got, 10)
		}
	})

	t.Run("Multiple push with same values", func(t *testing.T) {
		minStack := Constructor()

		minStack.Push(1)
		minStack.Push(1)
		minStack.Push(1)

		if got := minStack.GetMin(); got != 1 {
			t.Errorf("GetMin() = %v, want %v", got, 1)
		}

		minStack.Pop()

		if got := minStack.GetMin(); got != 1 {
			t.Errorf("GetMin() = %v, want %v", got, 1)
		}
	})

	t.Run("Push after multiple pops", func(t *testing.T) {
		minStack := Constructor()

		minStack.Push(4)
		minStack.Push(2)
		minStack.Push(1)
		minStack.Push(3)

		minStack.Pop()
		minStack.Pop()
		minStack.Pop()

		if got := minStack.Top(); got != 4 {
			t.Errorf("Top() = %v, want %v", got, 4)
		}

		minStack.Push(8)

		if got := minStack.GetMin(); got != 4 {
			t.Errorf("GetMin() = %v, want %v", got, 4)
		}
	})
}
