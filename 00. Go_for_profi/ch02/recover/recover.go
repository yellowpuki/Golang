package main

import "fmt"

func f1() {
	fmt.Println("Start f1")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover in f1")
		}
	}()

	f2()

	fmt.Println("End f1")
}

func f2() {
	fmt.Println("Start f2")

	panic("Panic in f2")

	fmt.Println("End f2")
}

func main() {
	fmt.Println("Start main")

	f1()

	fmt.Println("End main")
}
