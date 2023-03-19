package main

import (
	"fmt"

	"github.com/yellowpuki/randomSelect/internal/permutation"
)

func main() {
	var (
		selectSize int
		arraySize  int
	)

	fmt.Print("Enter size of select: ")
	fmt.Scan(&selectSize)
	fmt.Print("Enter size of array: ")
	fmt.Scan(&arraySize)

	p := permutation.NewPermutation(arraySize, selectSize)

	p.Generate()
	p.Print()

	fmt.Print("Enter another size of select: ")
	fmt.Scan(&selectSize)
	if err := p.SetSelectSize(selectSize); err != nil {
		fmt.Printf("Error set select size: %v\n", err)
		return
	}

	p.Generate()
	p.Print()
}
