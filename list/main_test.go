package list

import (
	"testing"
)

func TestNewList(t *testing.T) {
	l := New[int]()

	if actualValue := l.Len(); actualValue != 0 {
		t.Errorf("new list length is %d, expected %d", actualValue, 0)
	}
}

func TestAppend(t *testing.T) {
	l := New[int]()

	l.Append(1)
	l.Append(2)
	l.Append(3)

	if actualValue := l.Len(); actualValue != 3 {
		t.Errorf("list length is %d, expected %d", actualValue, 3)
	}

	if actualValue, ok := l.Get(0); !ok || actualValue != 1 {
		t.Errorf("list[0] is %d, expected %d", actualValue, 1)
	}
	if actualValue, ok := l.Get(1); !ok || actualValue != 2 {
		t.Errorf("list[1] is %d, expected %d", actualValue, 2)
	}
	if actualValue, ok := l.Get(2); !ok || actualValue != 3 {
		t.Errorf("list[2] is %d, expected %d", actualValue, 3)
	}
	if _, ok := l.Get(3); ok {
		t.Errorf("list[3] should not exist")
	}
}

func TestRemove(t *testing.T) {
	l := New[int]()

	l.Append(1)
	l.Append(2)
	l.Append(3)

	if actualValue := l.Len(); actualValue != 3 {
		t.Errorf("list length is %d, expected %d", actualValue, 3)
	}

	if ok := l.Remove(1); !ok {
		t.Errorf("expected to remove element at index 1")
	}

	if actualValue := l.Len(); actualValue != 2 {
		t.Errorf("list length is %d, expected %d", actualValue, 2)
	}

	if actualValue, ok := l.Get(1); !ok || actualValue != 3 {
		t.Errorf("list[1] is %d, expected %d", actualValue, 3)
	}

	l.Remove(0)
	if actualValue := l.Len(); actualValue != 1 {
		t.Errorf("list length is %d, expected %d", actualValue, 1)
	}

	l.Remove(1) // no effect
	if actualValue, ok := l.Get(0); actualValue != 3 || !ok {
		t.Errorf("list[0] is %d, expected %d", actualValue, 3)
	}

	l.Remove(0)
	if actualValue := l.IsEmpty(); actualValue != true {
		t.Errorf("list length is %t, expected %t", actualValue, true)
	}
}

func TestRemoveFirstNElements(t *testing.T) {
	l := New[int]()

	l.Append(1)
	l.Append(2)
	l.Append(3)

	if actualValue := l.Len(); actualValue != 3 {
		t.Errorf("list length is %d, expected %d", actualValue, 3)
	}

	if ok := l.RemoveFirstNElements(2); !ok {
		t.Errorf("expected to remove first 2 elements")
	}

	if actualValue := l.Len(); actualValue != 1 {
		t.Errorf("list length is %d, expected %d", actualValue, 1)
	}

	if actualValue, ok := l.Get(0); !ok || actualValue != 3 {
		t.Errorf("list[0] is %d, expected %d", actualValue, 3)
	}

	l.RemoveFirstNElements(1)
	if actualValue := l.IsEmpty(); actualValue != true {
		t.Errorf("list length is %t, expected %t", actualValue, true)
	}

	if ok := l.RemoveFirstNElements(1); ok {
		t.Errorf("expected to not remove any elements")
	}
}
