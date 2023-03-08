package main

import (
	"crypto/sha256"
	"fmt"
)

func getCountBytes(s1, s2 [32]byte) (n int) {
	for i, v := range s1 {
		if s2[i] != v {
			n++
		}
	}
	return n
}

func sha256Sum(s string) [32]byte {
	return sha256.Sum256([]byte(s))
}

func main() {
	c1 := sha256Sum("X")
	c2 := sha256Sum("x")

	fmt.Printf("%x\n%x\n%t\n%d\n", c1, c2, c1 == c2, getCountBytes(c1, c2))
}
