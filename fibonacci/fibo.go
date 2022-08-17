package main

import (
	"fmt"
)

func fibonacci(n int) int {
	first, second := 1, 1
	for i := 1; i < n; i++ {
		first, second = second, first+second
	}
	return first
}

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
	var (
		f1, f2, nom, flag = 0, 1, 1, 0
		t, fib            int
	)
	fmt.Scan(&fib)
	for flag != 1 {
		t, f1, f2 = f1+f2, f2, t
		nom += 1
		if f2 == fib {
			fmt.Printf("%d\n", nom)
			flag = 1
		} else if f2 > fib {
			fmt.Print(-1, "\n")
			flag = 1
		}
	}
}
