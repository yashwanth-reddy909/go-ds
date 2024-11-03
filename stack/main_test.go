package stack

import (
	"testing"
)

func TestNewStack(t *testing.T) {
	s := New[int]()
	if actualValue := s.IsEmpty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
}

func TestStackPush(t *testing.T) {
	s := New[int]()
	
	s.Push(1)
	s.Push(2)
	s.Push(3)
	
	if actualValue := s.Len(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, ok := s.Top(); actualValue != 3 || ok != true {
		t.Errorf("Got %v, %v expected %v, %v", actualValue, ok, 3, true)
	}
}

func TestStackPop(t *testing.T) {
	s := New[int]()

	s.Push(1)
	s.Push(2)

	if actualValue := s.Pop(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue, ok := s.Top(); actualValue != 1 || ok != true {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}
	if actualValue := s.Len(); actualValue != 1 {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}

	if actualValue := s.Pop(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := s.IsEmpty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
}