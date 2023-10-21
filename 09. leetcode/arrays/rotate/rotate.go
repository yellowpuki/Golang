package main

import "fmt"

func rotate(nums []int, k int) []int {
	res := make([]int, 0, len(nums))
	for i := len(nums) - k; i < len(nums); i++ {
		res = append(res, nums[i])
	}
	copy(res, nums[:len(nums)-k-1])
	return res
}

func shiftArrayRight(a []int, k int) {
	for k > 0 {
		n := len(a)
		tmp := a[n-1]
		for i := n - 1; i > 0; i-- {
			a[i] = a[i-1]
		}
		a[0] = tmp
		k--
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	shiftArrayRight(a, 3)
	fmt.Println(a)
}
