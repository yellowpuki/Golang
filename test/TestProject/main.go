package main

import (
	"fmt"
	"math"
	"strings"
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
	// a := "12"
	// //c := 100
	// b := []byte(a)
	// a += string(b)
	// fmt.Println(a)

	s := "How can mirrors be real if our eyes aren't real"
	fmt.Println(Rewerse(s))
}

func Rewerse(str string) string {
	word := ""
	res := ""
	for _, ch := range str {
		if ch == ' ' {
			res += word + " "
			word = ""
		} else {
			word = string(ch) + word
		}
	}
	return res + word
}

func isPowerOfThree(n int) bool {
	const (
		BASE = 3
		POW  = 19
	)

	return n > 0 && int(math.Pow(BASE, POW))%n == 0
}

func printShapeArea(s Shape) {
	fmt.Printf("Площадь фигуры: %.2f см\n", s.Area())
}

func ReverseWords(str string) string {
	var rev string
	var word string

	for _, i := range str {
		if i == ' ' {
			rev = rev + word + " " // Adds word and space to result
			word = ""              // Empties word variable
		} else {
			word = string(i) + word // Adds letter to temporary word variable
		}
	}

	return rev + word // reverse those words
}
func reverse(r []rune) []rune {
	for i, j := 0, len(r)-1; i < j; {
		r[i], r[j] = r[j], r[i]
		i++
		j--
	}
	return r
}

func ToJadenCase(str string) string {

	return strings.Title(str) // your code here...
}
