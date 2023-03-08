package test

import (
	"math/rand"
)

const selectBase = "SELECT * FROM user WHERE %s "

var refStringSlice = []string{
	" FIRST_NAME = 'Jack' ",
	" INSURANCE_NO = 333444555 ",
	" EFFECTIVE_FROM = SYSDATE "}

func FillArr(a []int) {
	for i := 0; i < len(a); i++ {
		a[i] = i * rand.Intn(1000)
	}
}

func Fibonacci(n int) int {
	var first, second = 1, 1

	for i := 1; i <= n; i++ {
		first, second = second, first+second
	}
	return first
}

// -----------------------------
// recSum ...
// -----------------------------
func RecSum(a []int) int {
	if len(a) == 0 {
		return 0
	}

	return a[0] + RecSum(a[1:])
}

// -----------------------------
// countEl ...
// -----------------------------
func CountEl(a []int) int {
	if len(a) == 0 {
		return 0
	}

	return 1 + CountEl(a[1:])
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
func MaxEl(a []int) int {
	if len(a) == 1 {
		return a[0]
	}

	return max(a[len(a)-1], MaxEl(a[:len(a)-1]))
}

// -----------------------------
// binSearchRec ...
// -----------------------------
func BinSearchRec(a []int, low int, hi int, target int) int {
	if low > hi {
		return -1
	}
	mid := (hi + low) / 2

	if target == a[mid] {
		return mid
	}

	if target < a[mid] {
		return BinSearchRec(a, low, mid-1, target)
	}

	return BinSearchRec(a, mid+1, hi, target)
}

// -----------------------------
// quickSort basic...
// -----------------------------
func QuickSortBase(a []int) []int {
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

	return append(append(QuickSortBase(less), pivot), QuickSortBase(greater)...)
}

// -----------------------------
// quickSort random pivot...
// -----------------------------
func QuickSortRnd(a []int) []int {
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

	position := rand.Int() % len(a)
	pivot := a[position]
	var less, greater []int

	for i := range a {
		if a[i] < pivot {
			less = append(less, a[i])
		}
		if a[i] > pivot {
			greater = append(greater, a[i])
		}
	}

	return append(append(QuickSortRnd(less), pivot), QuickSortRnd(greater)...)
}

// -----------------------------
// quickSort mid pivot...
// -----------------------------
func QuickSortMid(a []int) []int {
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
	left, right := 0, len(a)-1

	mid := (left + right) / 2
	pivot := a[mid]
	var less, greater []int

	for i := range a {
		if a[i] < pivot {
			less = append(less, a[i])
		}
		if a[i] > pivot {
			greater = append(greater, a[i])
		}
	}

	return append(append(QuickSortMid(less), pivot), QuickSortMid(greater)...)
}
