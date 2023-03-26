package main

import (
	"bytes"
	"errors"
	"fmt"
)

type Node[T comparable] struct {
	value T
	next  *Node[T]
}

type Stack[T comparable] struct {
	head *Node[T]
	len  int
}

func NewStack[T comparable]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) IsEmpty() bool {
	return s.head == nil
}

func (s *Stack[T]) Push(value T) {
	s.head = &Node[T]{
		value: value,
		next:  s.head,
	}
	s.len++
}

func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var empty T
		return empty, errors.New("stack is empty")
	}

	oldHead := s.head
	s.head = oldHead.next

	return oldHead.value, nil
}

func (s *Stack[T]) String() string {
	if s.IsEmpty() {
		return fmt.Sprintf("%T is empty", s)
	}

	var buffer bytes.Buffer
	sh := s.head
	for sh != nil {
		buffer.WriteString(fmt.Sprintf("%v\n-\n", sh.value))
		sh = sh.next
	}
	return buffer.String()
}

func main() {
	s := NewStack[int]()
	fmt.Println(s)

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	fmt.Println(s)

	s.Pop()

	fmt.Println(s)
}
