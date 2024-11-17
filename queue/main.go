package queue

import (
	"github.com/yashwanth-reddy909/go-ds/list"
)

type Queue[T any] struct {
	list *list.List[T]
}

// New returns a new intialized queue
func New[T any]() Queue[T] {
	return Queue[T]{
		list: list.New[T](),
	}
}

// Push adds an element to the end of the queue
func (q *Queue[T]) Push(e T) {
	q.list.Append(e)
}

// Pop removes the first element from the queue and returns it
func (q *Queue[T]) Pop() bool {
	if q.IsEmpty() {
		return false
	}

	q.list.RemoveFirstNElements(1)
	return true
}

// Peek returns the first element from the queue
func (q *Queue[T]) Peek() (T, bool) {
	if q.IsEmpty() {
		var t T
		return t, false
	}

	return q.list.Get(0)
}

// Len returns the number of elements in the queue
func (q *Queue[T]) Len() int {
	return q.list.Len()
}

// IsEmpty returns true if the queue is empty, false otherwise
func (q *Queue[T]) IsEmpty() bool {
	return q.list.IsEmpty()
}

// Clear removes all the elements from the queue
func (q *Queue[T]) Clear() {
	q.list.Clear()
}
