package ds

import "errors"

// Stack is a LIFO data structure backed by a slice.
type Stack struct {
	data []int
}

// NewStack creates a new empty Stack.
func NewStack() *Stack {
	return &Stack{data: make([]int, 0)}
}

// Push adds a value to the top of the stack.
func (s *Stack) Push(val int) {
	s.data = append(s.data, val)
}

// Pop removes and returns the value at the top of the stack.
// Returns an error if the stack is empty.
func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("pop from empty stack")
	}
	idx := len(s.data) - 1
	val := s.data[idx]
	s.data = s.data[:idx]
	return val, nil
}

// Peek returns the value at the top of the stack without removing it.
// Returns an error if the stack is empty.
func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("peek from empty stack")
	}
	return s.data[len(s.data)-1], nil
}

// IsEmpty returns true if the stack contains no elements.
func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

// Size returns the number of elements in the stack.
func (s *Stack) Size() int {
	return len(s.data)
}
