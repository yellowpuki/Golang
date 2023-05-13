package main

import "fmt"

func logger(prefix string) func(string) {
	return func(s string) {
		fmt.Printf("[%s]\t: %s\n", prefix, s)
	}
}

func main() {
	info := logger("INFO")
	warn := logger("WARN")
	err := logger("ERROR")

	info("information level")
	warn("warning level")
	err("error level")
}
