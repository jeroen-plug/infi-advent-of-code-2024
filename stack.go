package main

import (
	"container/list"
)

type Stack[T any] struct {
	list *list.List
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{list.New()}
}

func (s Stack[T]) Push(v T) {
	s.list.PushBack(v)
}

func (s Stack[T]) Pop() T {
	back := s.list.Back()
	v := back.Value.(T)
	s.list.Remove(back)
	return v
}

func (s Stack[_]) Len() int {
	return s.list.Len()
}
