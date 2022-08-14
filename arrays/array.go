package main

import "fmt"

func main() {
	var (
		workArray [10]byte
	)

	for idx := range workArray {
		fmt.Scan(&workArray[idx])
	}
	var x, y byte
	for i := 1; i <= 3; i++ {
		fmt.Scan(&x, &y)
		workArray[x], workArray[y] = workArray[y], workArray[x]
	}
	for idx := range workArray {
		fmt.Printf("%d ", workArray[idx])
	}
}
