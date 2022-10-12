package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := 0
	// simple for
	fmt.Println("1. Simple for loop")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}

	// while style loop
	fmt.Println("\n2.1 While style loop first version")
	for n < 10 {
		fmt.Printf("%d ", n)
		n++
	}

	// do While style loop
	fmt.Println("\n3.1 do While style loop first version")
	n = 0
	for ok := true; ok; ok = n < 10 {
		fmt.Printf("%d ", n)
		n++
	}

	fmt.Println("\n3.2 do While style loop second version")
	n = 0
	for {
		fmt.Printf("%d ", n)
		n++
		if n > 10 {
			break
		}
	}

	fmt.Println("\n3.3 do While style loop")
	n = 0
	for condition := false; !condition; {
		fmt.Printf("%d ", n)
		n++

		if n == 10 {
			condition = true
		}
	}

	// for range style loop
	fmt.Println("\n4. For range style loop")
	for n = range a {
		fmt.Printf("%d ", n)
	}
}
