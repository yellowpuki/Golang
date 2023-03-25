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
		return fmt.Sprintf("%T is empty", ss)
	}
	str := ""
	for _, v := range ss.st {
		str += fmt.Sprintf("%v ", v)
	}
	return fmt.Sprintf("%s<-Head", str)
}

func main() {
	stI := NewSliceStack[int]()
	stS := NewSliceStack[string]()

	fmt.Println(stI)
	fmt.Println(stS)

	stI.Push(1)
	stS.Push("Hello")
	stI.Push(2)
	stS.Push("World")
	stI.Push(3)
	stS.Push("!")
	fmt.Println(stI)
	fmt.Println(stS)

	for !stI.IsEmpty() && !stS.IsEmpty() {
		stI.Pop()
		stS.Pop()
		fmt.Println(stI)
		fmt.Println(stS)
	}
}
