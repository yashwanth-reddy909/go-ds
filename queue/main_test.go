package queue

import (
	"testing"
)

func TestNewQueue(t *testing.T) {
	q := New[int]()
	if actualValue := q.IsEmpty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
}

func TestQueuePush(t *testing.T) {
	q := New[int]()

	q.Push(1)
	q.Push(2)
	q.Push(3)

	if actualValue := q.Len(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, ok := q.Front(); actualValue != 1 || ok != true {
		t.Errorf("Got %v, %v expected %v, %v", actualValue, ok, 1, true)
	}
}

func TestQueuePop(t *testing.T) {
	q := New[int]()

	q.Push(1)
	q.Push(2)

	if actualValue := q.Pop(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	if actualValue := q.Len(); actualValue != 1 {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}
	if actualValue, ok := q.Front(); actualValue != 2 || ok != true {
		t.Errorf("Got %v, %v expected %v, %v", actualValue, ok, 2, true)
	}

	if actualValue := q.Pop(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := q.IsEmpty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	if actualValue := q.Pop(); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
}
