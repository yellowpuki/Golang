package main

import (
	"fmt"
)

// fibonacci ...
func fibonacci(n int) int {
	if n < 2 {
		return n
	}

	first, second := 1, 1
	for i := 1; i < n; i++ {
		first, second = second, first+second
	}

	return first
}

// fiboClosure ...
func fiboClosure() func() int {
	first, second := 0, 1
	return func() int {
		f := first
		first, second = second, first+second

		return f
	}
}

// fibo ...
func fibo(n int) int {
	if n < 1 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return fibo(n-1) + fibo(n-2)
}

func main() {

	f := fiboClosure()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", f())
	}
	// var (
	// 	f1, f2, nom, flag = 0, 1, 1, 0
	// 	t, fib            int
	// )
	// fmt.Scan(&fib)
	// for flag != 1 {
	// 	t, f1, f2 = f1+f2, f2, t
	// 	nom += 1
	// 	if f2 == fib {
	// 		fmt.Printf("%d", nom)
	// 		flag = 1
	// 	} else if f2 > fib {
	// 		fmt.Print(-1)
	// 		flag = 1
	// 	}
	// }
}
