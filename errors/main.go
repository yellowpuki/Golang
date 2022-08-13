package main

import (
	"fmt"
)

func divide(a int, b int) int {
	if b == 0 {
		panic("Divide by zero")
	}
	return a / b
}

func main() {
	var input int
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("check input data")
	} else {
		fmt.Println(divide(input, 5))
	}
}
