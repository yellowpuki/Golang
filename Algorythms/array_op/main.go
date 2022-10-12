package main

import (
	"fmt"
)

func printArray(a []int) {
	for i := 0; i < len(a); i++ {
		fmt.Printf("%3d ", a[i])
	}
	fmt.Printf("\n")
}

func copyArray(a []int) []int {
	b := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		b[i] = a[i]
	}
	return b
}

func reverseArray(a []int) {
	n := len(a)
	for i := 0; i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

func shiftArrayLeft(a []int) {
	n := len(a)
	tmp := a[0]
	for i := 0; i < n-1; i++ {
		a[i] = a[i+1]
	}
	a[n-1] = tmp
}

func shiftArrayRight(a []int) {
	n := len(a)
	tmp := a[n-1]
	for i := n - 1; i > 0; i-- {
		a[i] = a[i-1]
	}
	a[0] = tmp
}

func main() {
	a := []int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90}

	printArray(a)

	reverseArray(a)
	printArray(a)
	shiftArrayLeft(a)
	printArray(a)
	shiftArrayRight(a)
	printArray(a)
	b := copyArray(a)
	printArray(b)
}
