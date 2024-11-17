package heap

import (
	"cmp"
	"math"

	"github.com/yashwanth-reddy909/go-ds/list"
)

type Comparator[T any] func(x, y T) int

type Heap[T any] struct {
	list *list.List[T]
	cmp  Comparator[T]
}

// New returns a new heap with default comparator
func New[T cmp.Ordered]() *Heap[T] {
	return &Heap[T]{
		list: list.New[T](),
		cmp:  cmp.Compare[T],
	}
}

// NewWithComparator returns a new heap with the given comparator
func NewWithComparator[T any](cmp Comparator[T]) *Heap[T] {
	return &Heap[T]{
		list: list.New[T](),
		cmp:  cmp,
	}
}

// getParent returns the parent index with true if the parent exists
// otherwise it returns false
func getParent(index int) (int, bool) {
	curLevel, n := 0, 0
	for index+1 > n {
		n += int(math.Pow(2, float64(curLevel)))
		curLevel++
	}

	if curLevel == 1 {
		return curLevel, false
	}

	curLevelMaxElements := int(math.Pow(2, float64(curLevel-1)))
	parentLevelMaxElements := int(math.Pow(2, float64(curLevel-2)))
	curRowPosition := curLevelMaxElements - (n - index) + 1
	parentRowPosition := ((curRowPosition - 1) / 2) + 1

	return n - curLevelMaxElements - parentLevelMaxElements + parentRowPosition - 1, true
}

// returns the indexes of the children
func getChildren(index int) (int, int) {
	curLevel, n := 0, 0
	for index+1 > n {
		n += int(math.Pow(2, float64(curLevel)))
		curLevel++
	}

	curLevelMaxElements := int(math.Pow(2, float64(curLevel-1)))
	curRowPosition := curLevelMaxElements - (n - index) + 1
	return n + (2 * (curRowPosition - 1)), n + (2 * (curRowPosition - 1)) + 1
}

func (h *Heap[T]) bubbleUp(index int) {
	parentIndex, parentExists := getParent(index)
	if !parentExists {
		return
	}

	parent, _ := h.list.Get(parentIndex)
	e, _ := h.list.Get(index)
	if h.cmp(parent, e) == +1 {
		h.list.Swap(parentIndex, index)
		h.bubbleUp(parentIndex)
	}
}

// Push inserts the element into the heap
func (h *Heap[T]) Push(e T) {
	h.list.Append(e)
	h.bubbleUp(h.list.Len() - 1)
}

func (h *Heap[T]) bubbleDown(index int) {
	lIdx, rIdx := getChildren(index)
	l, lBool := h.list.Get(lIdx)
	r, rBool := h.list.Get(rIdx)
	cur, _ := h.list.Get(index)

	if ((lBool && rBool && h.cmp(l, r) == -1) || lBool) && h.cmp(l, cur) == -1 {
		_ = h.list.Swap(lIdx, index)
		h.bubbleDown(lIdx)
	} else if rBool && h.cmp(r, cur) == -1 {
		_ = h.list.Swap(rIdx, index)
		h.bubbleDown(rIdx)
	} else {
		return
	}
}

// Pop removes the root element
//
// if the heap is empty then it returns false
func (h *Heap[T]) Pop() bool {
	if h.IsEmpty() {
		return false
	}

	h.list.Swap(0, h.list.Len()-1)
	_ = h.list.Remove(h.list.Len() - 1)
	h.bubbleDown(0)
	return true
}

// Peek returns the root element
//
// if the heap is empty then it returns false
func (h *Heap[T]) Peek() (T, bool) {
	if h.IsEmpty() {
		var t T
		return t, false
	}
	return h.list.Get(0)
}

// Length return the length of the heap
func (h *Heap[T]) Len() int {
	return h.list.Len()
}

// IsEmpty returns true if the heap is empty, false otherwise
func (h *Heap[T]) IsEmpty() bool {
	return h.list.IsEmpty()
}

// Clear removes all the elements from the heap
func (h *Heap[T]) Clear() {
	h.list.Clear()
}
