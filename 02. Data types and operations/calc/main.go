package main

import (
	"fmt"
)

func clearCalc(a, b float32, op string) {
	a = 0
	b = 0
	op = ""
}

func add(a, b float32) (res float32, text string) {
	return a + b, "Сумма"
}

func sub(a, b float32) (res float32, text string) {
	return a - b, "Разность"
}

func mult(a, b float32) (res float32, text string) {
	return a * b, "Произведение"
}

func div(a, b float32) (res float32, text string) {
	return a / b, "Частное"
}

func main() {
	const oper = "+-*/"
	var (
		a, b float32
		res  float32
		op   string
		str  string
	)

	for {
		clearCalc(a, b, op)

		fmt.Print("Введите первое число: ")
		fmt.Scan(&a)

		fmt.Print("Выберите операнд (+, -, *, /): ")
		fmt.Scan(&op)

		fmt.Print("Введите второе число: ")
		fmt.Scan(&b)

		switch op {
		case string(oper[0]):
			{
				res, str = add(a, b)
			}
		case string(oper[1]):
			{
				res, str = sub(a, b)
			}
		case string(oper[2]):
			{
				res, str = mult(a, b)
			}
		case string(oper[3]):
			{
				if b == 0 {
					fmt.Println("Деление на 0 запрещено")
					clearCalc(a, b, op)
					continue
				} else {
					res, str = div(a, b)
				}
			}
		default:
			fmt.Println("Введен некорректный операнд")
			clearCalc(a, b, op)
			continue

		}
		fmt.Printf("\n%s чисел: %f\n\n", str, res)
	}
}
