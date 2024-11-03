package list

import (
	"slices"
)

const shrinkFactor = float32(0.25) // 25% of the capacity

type List[T any] struct {
	elements []T
}

func New[T any]() *List[T] {
	return &List[T]{}
}

func (q *List[T]) Append(e T) {
	q.elements = append(q.elements, e)
}

func (q *List[T]) Get(i int) (T, bool) {
	if i < 0 || i >= len(q.elements) {
		var t T
		return t, false
	}
	return q.elements[i], true
}

// shrink shrinks the list if the ratio of length to capacity is less than shrinkFactor
func (q *List[T]) shrink() {
	if cap(q.elements) != 0 && float32(len(q.elements)/cap(q.elements)) <= shrinkFactor {
		newList := make([]T, len(q.elements))
		copy(newList, q.elements)
		q.elements = newList
	}
}

// Remove removes the element at the given index
// returns true if the element was removed, false otherwise
//
// time complexity of this function is O(len(l)-index)
// sometimes O(n), if shrink happens
func (q *List[T]) Remove(index int) bool {
	if index < 0 || index >= len(q.elements) {
		return false
	}

	q.elements = slices.Delete(q.elements, index, index+1)
	q.shrink()
	return true
}

// RemoveFirstNElements removes the first n elements from the list
func (q *List[T]) RemoveFirstNElements(n int) bool {
	if len(q.elements) < n {
		return false
	}

	q.elements = q.elements[n:]
	q.shrink()
	return true
}

func (q *List[T]) Len() int {
	return len(q.elements)
}

func (q *List[T]) IsEmpty() bool {
	return q.Len() == 0
}
