package main

import (
	"factors/factors"
	"fmt"
)

func main() {

	var n int

	fmt.Print("Enter the number: ")
	fmt.Scan(&n)
	fmt.Printf("%v", factors.NumberFactors(n))
}
