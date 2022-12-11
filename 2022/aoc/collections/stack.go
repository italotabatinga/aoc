package collections

import (
	"fmt"
	"strings"
)

type Stack[T any] struct {
	elements []T
}

func NewStack[T any]() Stack[T] {
	stack := Stack[T]{}
	stack.elements = make([]T, 0)

	return stack
}

func (s *Stack[T]) Push(elem T) {
	s.elements = append(s.elements, elem)
}

func (s *Stack[T]) MultPush(elems ...T) {
	s.elements = append(s.elements, elems...)
}

func (s *Stack[T]) Pop() T {
	elem := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]

	return elem
}

func (s *Stack[T]) MultPop(n int) []T {
	elems := s.elements[len(s.elements)-n:]
	s.elements = s.elements[:len(s.elements)-n]

	return elems
}

func (s Stack[T]) Top() T {
	return s.elements[len(s.elements)-1]
}

func (s Stack[T]) Len() int {
	return len(s.elements)
}

func (s Stack[T]) String() string {
	var sb strings.Builder
	sb.WriteString("Stack(")
	for i, r := range s.elements {
		sb.WriteString(fmt.Sprintf("%v", r))

		if i != len(s.elements)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString(")")
	return sb.String()
}
