package collections

import (
	"fmt"
	"strings"
)

type Queue[T any] struct {
	elements []T
}

func NewQueue[T any]() Queue[T] {
	queue := Queue[T]{}
	queue.elements = make([]T, 0)

	return queue
}

func (q Queue[T]) Index(index int) T {
	return q.elements[index]
}

func (q *Queue[T]) Push(elem T) {
	q.elements = append(q.elements, elem)
}

func (q *Queue[T]) MultPush(elems ...T) {
	q.elements = append(q.elements, elems...)
}

func (q *Queue[T]) Pop() T {
	elem := q.elements[0]
	q.elements = q.elements[1:]

	return elem
}

func (q *Queue[T]) MultPop(n int) []T {
	elems := q.elements[:n-1]
	q.elements = q.elements[n:]

	return elems
}

func (q Queue[T]) Front() T {
	return q.elements[0]
}

func (q Queue[T]) Rear() T {
	return q.elements[len(q.elements)-1]
}

func (q Queue[T]) Len() int {
	return len(q.elements)
}

func (q Queue[T]) String() string {
	var sb strings.Builder
	sb.WriteString("Queue(")
	for i, r := range q.elements {
		sb.WriteString(fmt.Sprintf("%v", r))

		if i != len(q.elements)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString(")")
	return sb.String()
}
