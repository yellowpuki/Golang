package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("You are using %s compiler on %s pc\n", runtime.Compiler, runtime.GOARCH)
	fmt.Printf("Using Go version: %s\n", runtime.Version())
	fmt.Printf("Number of CPU: %d\n", runtime.NumCPU())
	fmt.Printf("Number of goroutines: %d\n", runtime.NumGoroutine())
}
