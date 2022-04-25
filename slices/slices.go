package main

import (
	"fmt"
)

func main() {
	baseArray := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Base array: %v\n", baseArray)

	baseSlice := baseArray[1:4]
	fmt.Printf(
		"Slice based on Base array length %d and capacity %d: %v\n",
		len(baseSlice),
		cap(baseSlice),
		baseSlice,
	)

	pointer := fmt.Sprintf("%p", baseSlice)

	baseSlice = append(baseSlice, 10)
	fmt.Printf("Array: %v\n", baseArray)
	fmt.Printf(
		"Slice length %d and capacity %d: %v\n",
		len(baseSlice),
		cap(baseSlice),
		baseSlice,
	)
	fmt.Println(pointer == fmt.Sprintf("%p", baseSlice))

	baseSlice = append(baseSlice, 11, 12, 13, 14, 15)
	fmt.Printf("Array: %v\n", baseArray)
	fmt.Printf(
		"Slice length %d and capacity %d: %v\n",
		len(baseSlice),
		cap(baseSlice),
		baseSlice,
	)
	fmt.Println(pointer == fmt.Sprintf("%p", baseSlice))

	/*	a := []int{1, 2, 3}
		b := make([]int, 3, 3)
		n := copy(b, a)

		fmt.Printf("a = %v\n", a)
		fmt.Printf("b = %v\n", b)
		fmt.Printf("Copied %d elements\n", n)*/

}
