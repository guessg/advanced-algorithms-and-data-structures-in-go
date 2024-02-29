package stack

// Stack represents a stack
type Stack[T any] struct {
	data []T
}

// NewStack creates a new stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Push pushes an element onto the stack
func (s *Stack[T]) Push(e T) {
	s.data = append(s.data, e)
}

// Pop pops an element from the stack
func (s *Stack[T]) Pop() (t T) {
	if len(s.data) == 0 {
		return
	}

	e := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return e
}

// Peek returns the top element of the stack
func (s *Stack[T]) Peek() (t T) {
	if len(s.data) == 0 {
		return
	}

	return s.data[len(s.data)-1]
}

// IsEmpty returns true if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}
