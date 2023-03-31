package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a1 []int
	a2 := []int{2}

	var m1 map[int]int
	m2 := map[int]int{1: 1}

	var f1 func() int
	f2 := func() int { return 2 }

	//данные которые можно сравнивать, а также структуры содержащие их в себе так же можно сравнивать
	type comparableData struct {
		num     int
		fp      float64
		complex complex64
		str     string
		char    rune
		yes     bool
		events  <-chan string
		handler interface{}
		ref     *byte
		raw     [10]byte
	}

	type uncomparableData struct {
		checks [10]func() bool
		doit   func() bool
		m      map[string]string
		bytes  []byte
	}

	//fmt.Println("Слайсы нельзя сравнивать между собой, только с nil", a1 == a2)
	fmt.Println("a1==nil?", a1 == nil)
	fmt.Println("a2==nil?", a2 == nil)

	//fmt.Println("Мапы тоже нельзя сравнивать между собой, только с nil", m1 == m2)
	fmt.Println("m1 == nil?", m1 == nil)
	fmt.Println("m2 == nil?", m2 == nil)

	// fmt.Println("Цункции также нальзя сравнивать между собой, только с nil", f1 == f2)
	fmt.Println("f1 == nil", f1 == nil)
	fmt.Println("f2 == nil", f2 == nil)
	//но можно использовать функцию из пакете reflect reflect.DeepEqual()
	v1 := comparableData{}
	v2 := comparableData{}
	fmt.Println("v1 == v2", v1 == v2)

	uv1 := uncomparableData{}
	uv2 := uncomparableData{}

	//fmt.Println("uv1 == uv2", uv1 == uv2)
	fmt.Println("uv1 == uv2", reflect.DeepEqual(uv1, uv2))
}
