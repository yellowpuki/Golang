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

	a := []int{12, 22, 3, 14, 5, 16, 7, -8, 9, 0}

	fmt.Println(recSum([]int{2, 4, 6}))
	fmt.Println(countEl([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}))
	fmt.Println(maxEl([]int{12, 2, 3, 4, 5, 6, 7, 8, 9, 0}))
	fmt.Println(binSearchRec(a, 0, len(a)-1, 2))
	fmt.Println(quickSort(a))

}

func fibonacci(n int) int {
	var first, second = 1, 1

	for i := 1; i <= n; i++ {
		first, second = second, first+second
	}
	return first
}

// -----------------------------
// recSum ...
// -----------------------------
func recSum(a []int) int {
	if len(a) == 0 {
		return 0
	}

	return a[0] + recSum(a[1:])
}

// -----------------------------
// countEl ...
// -----------------------------
func countEl(a []int) int {
	if len(a) == 0 {
		return 0
	}

	return 1 + countEl(a[1:])
}

// -----------------------------
// max ...
// -----------------------------
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// -----------------------------
// maxEl ...
// -----------------------------
func maxEl(a []int) int {
	if len(a) == 1 {
		return a[0]
	}

	return max(a[len(a)-1], maxEl(a[:len(a)-1]))
}

// -----------------------------
// binSearchRec ...
// -----------------------------
func binSearchRec(a []int, low int, hi int, target int) int {
	if low > hi {
		return -1
	}
	mid := (hi + low) / 2

	if target == a[mid] {
		return mid
	}

	if target < a[mid] {
		return binSearchRec(a, low, mid-1, target)
	}

	return binSearchRec(a, mid+1, hi, target)
}

// -----------------------------
// quickSort ...
// -----------------------------
func quickSort(a []int) []int {
	if len(a) < 2 {

		return a
	}

	if len(a) == 2 {
		left, right := 0, len(a)-1
		if a[left] > a[right] {
			a[left], a[right] = a[right], a[left]
		}

		return a
	}

	pivot := a[0]
	var less, greater []int

	for i := range a {
		if a[i] < pivot {
			less = append(less, a[i])
		}
		if a[i] > pivot {
			greater = append(greater, a[i])
		}
	}

	return append(append(quickSort(less), pivot), quickSort(greater)...)
}
