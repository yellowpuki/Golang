package main

import (
	"fmt"
)

type Calculator interface {
	add() float64
	sub() float64
	mul() float64
	div() float64
}

type Expression struct {
	value1    float64
	value2    float64
	operation string
}

func NewExpression(v1, v2, op interface{}) *Expression {
	return &Expression{
		value1:    iToFloat(v1),
		value2:    iToFloat(v2),
		operation: iToStr(op),
	}
}

func iToFloat(i interface{}) float64 {
	v, ok := i.(float64)
	if !ok {
		panic(fmt.Sprintf("value=%v: %T", i, i))
	}

	return v
}

func iToStr(i interface{}) string {
	v, ok := i.(string)
	if !ok {
		panic(fmt.Sprintf("value=%v: %T", i, i))
	}

	return v
}

func (e *Expression) Calculate() float64 {
	switch e.operation {
	case "+":
		return e.add()
	case "-":
		return e.sub()
	case "*":
		return e.mul()
	case "/":
		return e.div()
	default:
		panic(fmt.Sprintf("неизвестная операция"))
	}
}

func (e *Expression) add() float64 {

	return e.value1 + e.value2
}

func (e *Expression) sub() float64 {

	return e.value1 - e.value2
}

func (e *Expression) mul() float64 {

	return e.value1 * e.value2
}

func (e *Expression) div() float64 {

	if e.value2 == 0 {
		panic(fmt.Sprintf("деление на ноль"))
	}

	return e.value1 / e.value2
}

func main() {
	a := 1.0
	b := 2.2

	op := "+"

	// value1, err := itof64(a)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	//
	// value2, err := itof64(b)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	//
	// operator, err := iToString(op)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	//
	// calculate(value1, value2, operator)
	//

	ex := NewExpression(a, b, op)

	result := ex.Calculate()

	fmt.Printf("%.4f\n", result)
}

func errorValue(i interface{}) {
	fmt.Printf("value=%v:%T\n", i, i)
}

func itof64(i interface{}) (float64, error) {
	if v, ok := i.(float64); ok {
		return v, nil
	}

	return 0.0, fmt.Errorf("value=%v:%T", i, i)
}

func iToString(i interface{}) (string, error) {
	if v, ok := i.(string); ok {
		return v, nil
	}
	return "", fmt.Errorf("неизвестный оператор")
}

func calculate(v1, v2 float64, op string) {
	switch op {
	case "+":
		fmt.Printf("%.4f\n", v1+v2)
	case "-":
		fmt.Printf("%.4f\n", v1-v2)
	case "*":
		fmt.Printf("%.4f\n", v1*v2)
	case "/":
		if v2 == 0.0 {
			fmt.Printf("%.4f", 0.0)
		} else {
			fmt.Printf("%.4f\n", v1/v2)
		}
	default:
		fmt.Println("неизвестный оператор")
		return
	}
}
