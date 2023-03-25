package main

import (
	"errors"
	"fmt"
)

type SliceStack[T comparable] struct {
	st []T
}

func NewSliceStack[T comparable]() *SliceStack[T] {
	return &SliceStack[T]{
		st: make([]T, 0),
	}
}

func (ss *SliceStack[T]) IsEmpty() bool {
	return len(ss.st) == 0
}

func (ss *SliceStack[T]) Push(val T) {
	ss.st = append(ss.st, val)
}

func (ss *SliceStack[T]) Pop() (T, error) {
	if ss.IsEmpty() {
		var empty T
		return empty, errors.New("stack is empty")
	}

	head := len(ss.st)
	val := ss.st[head-1]
	ss.st = ss.st[:head-1]

	return val, nil
}

func (ss *SliceStack[T]) String() string {
	if ss.IsEmpty() {
		return "stack is empty"
	}
	str := ""
	for _, v := range ss.st {
		str += fmt.Sprintf("%v ", v)
	}
	return fmt.Sprintf("%s<-Head", str)
}

func main() {
	st := NewSliceStack[int]()

	fmt.Println(st)

	st.Push(1)
	st.Push(2)
	st.Push(3)
	fmt.Println(st)

	for !st.IsEmpty() {
		st.Pop()
		fmt.Println(st)
	}
}
