package heap

import (
	"testing"
)

func TestPush(t *testing.T) {
	heap := New[int]()
	heap.Push(1)
	heap.Push(2)
	heap.Push(3)

	if actualValue := heap.Len(); actualValue != 3 {
		t.Errorf("list length is %d, expected %d", actualValue, 3)
	}

	heap.Push(0)
	if actualValue := heap.Len(); actualValue != 4 {
		t.Errorf("list length is %d, expected %d", actualValue, 4)
	}
}

func TestPop(t *testing.T) {
	heap := New[int]()
	heap.Push(1)
	heap.Push(2)

	if actualValue := heap.Pop(); actualValue != true {
		t.Errorf("heap remove is %t, expected %t", actualValue, true)
	}
	if actualValue := heap.Len(); actualValue != 1 {
		t.Errorf("heap length is %d, expected %d", actualValue, 1)
	}

	heap.Push(0)
	if actualValue := heap.Pop(); actualValue != true {
		t.Errorf("heap remove is %t, expected %t", actualValue, true)
	}
	if actualValue := heap.Len(); actualValue != 1 {
		t.Errorf("heap length is %d, expected %d", actualValue, 0)
	}
	if actualValue := heap.Pop(); actualValue != true {
		t.Errorf("heap remove is %t, expected %t", actualValue, true)
	}
}

func TestPeek(t *testing.T) {
	heap := New[int]()
	heap.Push(1)
	heap.Push(2)
	heap.Push(0)

	if actualValue, ok := heap.Peek(); !ok || actualValue != 0 {
		t.Errorf("heap peek is %d, expected %d", actualValue, 0)
	}

	heap.Pop()
	if actualValue, ok := heap.Peek(); !ok || actualValue != 1 {
		t.Errorf("heap peek is %d, expected %d", actualValue, 1)
	}

	heap.Pop()
	heap.Push(-1)
	if actualValue, ok := heap.Peek(); !ok || actualValue != -1 {
		t.Errorf("heap peek is %d, expected %d", actualValue, -1)
	}
}
