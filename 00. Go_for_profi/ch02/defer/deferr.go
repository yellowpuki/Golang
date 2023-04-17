package main

import "fmt"

func d0() {
	for i := 3; i > 0; i-- {
		fmt.Print(i, " ")
	}
}

func d1() {
	for i := 3; i > 0; i-- {
		defer fmt.Print(i, " ")
	}
	fmt.Println()
}

func d2() {
	for i := 3; i > 0; i-- {
		defer func() {
			fmt.Print(i, " ")
		}()
	}
	fmt.Println()
}

func d3() {
	for i := 3; i > 0; i-- {
		defer func(n int) {
			fmt.Print(n, " ")
		}(i)
	}
	fmt.Println()
}

func d4() {
	for i := 3; i > 0; i-- {
		i := i
		defer func() {
			fmt.Print(i, " ")
		}()
	}
	fmt.Println()
}

func main() {
	d0()
	d1()
	d2()
	d3()
	d4()
}
