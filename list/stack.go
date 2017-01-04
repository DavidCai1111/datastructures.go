package list

import (
	"container/list"

	ds "github.com/DavidCai1993/datastructures.go"
)

// Stack represents a stack.
type Stack struct {
	l *list.List
}

// NewStack returns a new stack instance.
func NewStack() *Stack {
	return &Stack{l: list.New()}
}

// Push adds an element to the top of the stack.
func (s *Stack) Push(c ds.Comparable) {
	s.l.PushFront(c)
}

// Pop removes the top element from the stack.
func (s *Stack) Pop() ds.Comparable {
	return s.l.Remove(s.l.Front()).(ds.Comparable)
}

// Top returns the top element from the stack.
func (s *Stack) Top() ds.Comparable {
	return s.l.Front().Value.(ds.Comparable)
}

// Empty checks if the stack is empty or not.
func (s *Stack) Empty() bool {
	return s.l.Len() == 0
}
