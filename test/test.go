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
