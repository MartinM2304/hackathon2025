package models

import (
	"errors"
	"fmt"
)

type Queue[T any] struct {
	Elements []T
}

func (q *Queue[T]) Enqueue(elem T) {
	if q.Elements == nil {
		q.Elements = []T{}
	}
	q.Elements = append(q.Elements, elem)
}

func (q *Queue[T]) Dequeue() T {
	if q.IsEmpty() {
		fmt.Println("UnderFlow")
		return *new(T)
	}
	element := q.Elements[0]
	if q.GetLength() == 1 {
		q.Elements = nil
		return element
	}
	q.Elements = q.Elements[1:]
	return element
}

func (q *Queue[T]) GetLength() int {
	return len(q.Elements)
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.Elements) == 0
}

func (q *Queue[T]) Peek() (T, error) {
	if q.IsEmpty() {
		return *new(T), errors.New("empty queue")
	}
	return q.Elements[0], nil
}
