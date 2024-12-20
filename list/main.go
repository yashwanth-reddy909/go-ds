package list

import (
	"slices"
)

const shrinkFactor = float32(0.25) // 25% of the capacity

type List[T any] struct {
	elements []T
}

// New returns a new intialized list
func New[T any]() *List[T] {
	return &List[T]{}
}

// append adds an element to the end of the list
func (l *List[T]) Append(e T) {
	l.elements = append(l.elements, e)
}

// Get returns the element at the given index
// return ok false if the index is out of bounds
func (l *List[T]) Get(i int) (T, bool) {
	if i < 0 || i >= len(l.elements) {
		var t T
		return t, false
	}
	return l.elements[i], true
}

// shrink shrinks the list if the ratio of length to capacity is less than shrinkFactor
func (l *List[T]) shrink() {
	if cap(l.elements) != 0 && float32(len(l.elements)/cap(l.elements)) <= shrinkFactor {
		newList := make([]T, len(l.elements))
		copy(newList, l.elements)
		l.elements = newList
	}
}

// Remove removes the element at the given index
// returns true if the element was removed, false otherwise
//
// time complexity of this function is O(len(l)-index)
// sometimes O(n), if shrink happens
func (l *List[T]) Remove(index int) bool {
	if index < 0 || index >= len(l.elements) {
		return false
	}

	l.elements = slices.Delete(l.elements, index, index+1)
	l.shrink()
	return true
}

// RemoveFirstNElements removes the first n elements from the list
func (l *List[T]) RemoveFirstNElements(n int) bool {
	if len(l.elements) < n {
		return false
	}

	l.elements = l.elements[n:]
	l.shrink()
	return true
}

func (l *List[T]) Len() int {
	return len(l.elements)
}

func (l *List[T]) IsEmpty() bool {
	return l.Len() == 0
}

func (l *List[T]) Clear() {
	l.elements = []T{}
}

// Swap returns true if successfully swaped
// if i, j are out of indexed then returns false
func (l *List[T]) Swap(i, j int) bool {
	if i >= 0 && j >= 0 && i < l.Len() && j < l.Len() {
		l.elements[i], l.elements[j] = l.elements[j], l.elements[i]
		return true
	}
	return false
}

func (l *List[T]) Print() {
	for i := 0; i < l.Len(); i++ {
		print(l.elements[i], " ")
	}
	println()
}
