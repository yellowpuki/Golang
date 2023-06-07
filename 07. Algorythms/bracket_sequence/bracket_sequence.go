package main

import (
	"errors"
	"fmt"
	"strings"
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

func CheckSequence(input string) bool {
	stack := NewSliceStack[rune]()
	open := "{[("

	for _, r := range input {
		if strings.Contains(open, string(r)) {
			stack.Push(r)
		} else {
			if stack.IsEmpty() {
				return false
			}

			top, _ := stack.Pop()
			if top == '(' && r != ')' ||
				top == '{' && r != '}' ||
				top == '[' && r != ']' {
				return false
			}
		}
	}

	return stack.IsEmpty()
}

func main() {
	str := "{}"

	fmt.Printf("%+v", CheckSequence(str))
}
