package test

import "testing"

func BenchmarkQuickSortBase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := make([]int, 10000)
		FillArr(a)
		QuickSortBase(a)
	}
}

func BenchmarkQuickSortMid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := make([]int, 10000)
		FillArr(a)
		QuickSortMid(a)
	}
}

func BenchmarkQuickSortRnd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := make([]int, 10000)
		FillArr(a)
		QuickSortRnd(a)
	}
}
