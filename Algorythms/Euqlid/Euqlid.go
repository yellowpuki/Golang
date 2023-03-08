package main

import "fmt"

func euclidGcd(a int, b int) int {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}
	return a
}

func main() {
	var a, b int
	fmt.Println("Введите первое число:")
	fmt.Scan(&a)
	fmt.Println("введите второе число:")
	fmt.Scan(&b)
	fmt.Printf("NOD of %d and %d = %d\n", a, b, euclidGcd(a, b))
}
