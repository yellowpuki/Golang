package main

import (
	"fmt"
)

type Sorter interface {
	Sort(d array2d)
}

type array2d struct {
	array [][]int
	row   int
	col   int
}

func swap(a *int, b *int) {
	*a, *b = *b, *a
}

func (a *array2d) initArray2d() array2d {
	a.array = make([][]int, a.row)
	for i := range a.array {
		a.array[i] = make([]int, a.col)
	}
	return *a
}

func (a *array2d) fillArray2d() {
	for i := 0; i < a.row; i++ {
		for j := 0; j < a.col; j++ {
			a.array[i][j] = (i + 1) * (j + 1)
		}
	}
}

func (a *array2d) printArray2d() {
	for i := 0; i < a.row; i++ {
		for j := 0; j < a.col; j++ {
			fmt.Printf("%4d", a.array[i][j])
		}
		fmt.Printf("\n")
	}
}

func (a *array2d) bubleSort2d(direction int) {
	for i := 0; i < a.row; i++ {
		for j := 0; j < a.col; j++ {
			for k := 0; k < a.row; k++ {
				for n := 0; n < a.col; n++ {
					switch direction {
					case 0:
						{
							if a.array[i][j] < a.array[k][n] {
								swap(&a.array[i][j], &a.array[k][n])
							}
						}
					case 1:
						{
							if a.array[i][j] > a.array[k][n] {
								swap(&a.array[i][j], &a.array[k][n])
							}
						}
					}
				}
			}

		}
	}
}

func (a *array2d) findMax() (idxI int, idxJ int, max int) {
	var m int = a.array[0][0]
	for i := 0; i < a.row; i++ {
		for j := 0; j < a.col; j++ {
			if a.array[i][j] > m {
				m = a.array[i][j]
				max = m
				idxI = i
				idxJ = j
			}
		}
	}
	return max, idxI, idxJ
}

func main() {
	const (
		ASC  = 0
		DESC = 1
		ROW  = 10
		COL  = 10
	)

	a := array2d{
		array: [][]int{},
		row:   ROW,
		col:   COL,
	}

	a.initArray2d()
	a.fillArray2d()
	a.printArray2d()
	max, i, j := a.findMax()
	fmt.Printf("Максимальное число=%d;Индекс i=%d; Индекс j=%d\n", max, i, j)
	a.bubleSort2d(ASC)
	a.printArray2d()

}
