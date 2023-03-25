package main

import (
	"fmt"
)

func binarySearch(target []int, value int) int {
	start := 0
	end := len(target) - 1

	for start <= end {
		mid := (start + end) / 2
		if target[mid] == value {
			return mid
		}
		if target[mid] > value {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}

func main() {
	var (
		list = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		val  int
	)
	fmt.Println(list)
	fmt.Println("Input number to search position")
	fmt.Scan(&val)
	res := binarySearch(list, val)
	if res >= 0 {
		fmt.Println("Position: ", res)
	} else {
		fmt.Println("Not found value in list")
	}
}
