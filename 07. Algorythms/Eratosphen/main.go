package main

import "fmt"

// PrintSieve ...
func PrintSieve(arr []int) {
	for i := 2; i < len(arr); i++ {
		if arr[i] == 0 {
			fmt.Printf("%3d", i)
		}
	}
	fmt.Println()
}

// Sieve ...
func Sieve(arr []int) {
	for i := 2; i*i < len(arr); i++ {
		if arr[i] == 0 {
			for j := i * i; j < len(arr); j += i {
				arr[j] = 1
			}
		}
	}
}

func main() {
	a := make([]int, 100)

	Sieve(a)
	PrintSieve(a)
}
