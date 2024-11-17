package stack

import (
	"github.com/yashwanth-reddy909/go-ds/list"
)

type Stack[T any] struct {
	list *list.List[T]
}

// New returns a new intialized stack
func New[T any]() Stack[T] {
	return Stack[T]{
		list: list.New[T](),
	}
}

// Push adds an element to the end of the stack
func (s *Stack[T]) Push(e T) {
	s.list.Append(e)
}

// Pop removes the last element from the stack and returns it
func (s *Stack[T]) Pop() bool {
	if s.list.IsEmpty() {
		return false
	}

	s.list.Remove(s.list.Len() - 1)
	return true
}

// Peek returns the last element from the stack
func (s *Stack[T]) Peek() (T, bool) {
	if s.list.IsEmpty() {
		var t T
		return t, false
	}

	return s.list.Get(s.list.Len() - 1)
}

// Len returns the number of elements in the stack
func (s *Stack[T]) Len() int {
	return s.list.Len()
}

// IsEmpty returns true if the stack is empty, false otherwise
func (s *Stack[T]) IsEmpty() bool {
	return s.list.IsEmpty()
}

// Clear removes all the elements from the stack
func (s *Stack[T]) Clear() {
	s.list.Clear()
}
