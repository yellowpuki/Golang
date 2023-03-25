package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	go function()

	go func() {
		for i := 10; i < 20; i++ {
			fmt.Print(i, " ")
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Println()
}

func function() {
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
}

func worker(id int, limit <-chan time.Time, wg *sync.WaitGroup) {
	defer wg.Done()
	time := <-limit
	fmt.Printf("worker %d done work, time: %v\n", id, time)
}

func calculator1(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	ch := make(chan int)

	go func(ch chan int) {
		defer close(ch)
	}(ch)

	select {
	case x := <-firstChan:
		ch <- x * x
		return ch
	case x := <-secondChan:
		ch <- x * 3
		return ch
	case <-stopChan:

	}
	return ch
}

func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {
	res := make(chan int)

	go func(ch chan int) {
		defer close(ch)

		sum := 0
		for {
			select {
			case x := <-arguments:
				sum += x
			case <-done:
				res <- sum
				return
			}
		}

	}(res)

	return res
}
