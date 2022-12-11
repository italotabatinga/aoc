package collections

type Set[T comparable] struct {
	elements map[T]bool
	size     int
}

func (s *Set[T]) Add(x T) {
	if _, ok := s.elements[x]; !ok {
		s.size += 1
		s.elements[x] = true
	}
}

func (s *Set[T]) Remove(x T) {
	if _, ok := s.elements[x]; ok {
		s.size -= 1
		delete(s.elements, x)
	}
}

func (s Set[T]) Size() int {
	return s.size
}

func NewSet[T comparable]() Set[T] {
	return Set[T]{
		elements: make(map[T]bool, 0),
		size:     0,
	}
}
