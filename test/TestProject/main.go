package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float32
}

type Square struct {
	sideLength float32
}

func (s Square) Area() float32 {
	return s.sideLength * s.sideLength
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return float32(math.Pi) * c.radius
}

func main() {
	/*	Square := Square{10.5}
		circle := Circle{5.7}

		printShapeArea(Square)
		printShapeArea(circle)

		a := 101.0 / 2.0
		fmt.Print(a + float64(uint8(a)))
	*/
	a := "12"
	//c := 100
	b := []byte(a)
	a += string(b)
	fmt.Println(a)
}

func printShapeArea(s Shape) {
	fmt.Printf("Площадь фигуры: %.2f см\n", s.Area())
}
