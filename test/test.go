package main

import (
	"fmt"
	//"strings"
)

const selectBase = "SELECT * FROM user WHERE %s "

var refStringSlice = []string{
	" FIRST_NAME = 'Jack' ",
	" INSURANCE_NO = 333444555 ",
	" EFFECTIVE_FROM = SYSDATE "}

func main() {

	//	sentence := strings.Join(refStringSlice, "AND")
	//	fmt.Printf(selectBase+"\n", sentence)
	// for i := 0; i <= 10; i++ {
	// 	fmt.Printf("%d ", fibonacci(i))
	// }

	fmt.Println(recSum([]int{2, 4, 6}))
	fmt.Println(countEl([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}))
	fmt.Println(maxEl([]int{12, 2, 3, 4, 5, 6, 7, 8, 9, 0}))
}

func fibonacci(n int) int {
	var first, second = 1, 1

	for i := 1; i <= n; i++ {
		first, second = second, first+second
	}
	return first
}

func recSum(a []int) int {
	if len(a) == 0 {
		return 0
	}

	return a[0] + recSum(a[1:])
}

func countEl(a []int) int {
	if len(a) == 0 {
		return 0
	}

	return 1 + countEl(a[1:])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxEl(a []int) int {
	if len(a) == 1 {
		return a[0]
	}

	return max(a[len(a)-1], maxEl(a[:len(a)-1]))
}
