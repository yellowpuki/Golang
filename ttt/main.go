package main

import "fmt"

func bill(a int, b int, c int, d int) int {
	if d > b {
		return a + (d-b)*c
	}
	return a
}

func equalPie(n int) int {
	var (
		p int = 1
		e int = 0
	)
	for p < n {
		e += 1
		p *= 2
	}
	return e
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Printf("%d", equalPie(n))
}
