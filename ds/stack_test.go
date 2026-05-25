package ds_test

import (
	"DeepTest/ds"
	"testing"
)

func TestNewStack(t *testing.T) {
	s := ds.NewStack()
	if !s.IsEmpty() {
		t.Error("new stack should be empty")
	}
	if s.Size() != 0 {
		t.Errorf("new stack size should be 0, got %d", s.Size())
	}
}

func TestStackPush(t *testing.T) {
	s := ds.NewStack()
	s.Push(10)
	s.Push(20)

	if s.Size() != 2 {
		t.Errorf("size = %d, want 2", s.Size())
	}
	val, err := s.Peek()
	if err != nil {
		t.Fatalf("Peek() error: %v", err)
	}
	if val != 20 {
		t.Errorf("Peek() = %d, want 20", val)
	}
}

func TestStackPop(t *testing.T) {
	s := ds.NewStack()
	s.Push(10)
	s.Push(20)

	val, err := s.Pop()
	if err != nil {
		t.Fatalf("Pop() error: %v", err)
	}
	if val != 20 {
		t.Errorf("Pop() = %d, want 20", val)
	}
	if s.Size() != 1 {
		t.Errorf("size = %d, want 1", s.Size())
	}

	val, err = s.Pop()
	if err != nil {
		t.Fatalf("Pop() error: %v", err)
	}
	if val != 10 {
		t.Errorf("Pop() = %d, want 10", val)
	}
	if !s.IsEmpty() {
		t.Error("stack should be empty after popping all elements")
	}
}

func TestStackPopEmpty(t *testing.T) {
	s := ds.NewStack()
	_, err := s.Pop()
	if err == nil {
		t.Error("expected error when popping empty stack")
	}
}

func TestStackPeekEmpty(t *testing.T) {
	s := ds.NewStack()
	_, err := s.Peek()
	if err == nil {
		t.Error("expected error when peeking empty stack")
	}
}

func TestStackLIFO(t *testing.T) {
	s := ds.NewStack()
	values := []int{1, 2, 3, 4, 5}
	for _, v := range values {
		s.Push(v)
	}
	for i := len(values) - 1; i >= 0; i-- {
		val, err := s.Pop()
		if err != nil {
			t.Fatalf("Pop() error: %v", err)
		}
		if val != values[i] {
			t.Errorf("Pop() = %d, want %d", val, values[i])
		}
	}
}
