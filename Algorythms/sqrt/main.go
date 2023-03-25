package main

import (
	"fmt"
	"math"
)

const round = 10000000000000000.0

type ErrNegativeSqrt float64

// Error ...
func (e ErrNegativeSqrt) Error() string {

	return fmt.Sprintf("cannot Sqrt negative number: %g", float64(e))
}

// Sqrt ...
func Sqrt(x float64) (float64, error) {

	if x < 0 {
		return 0.0, ErrNegativeSqrt(x)
	}
	guess := 1.0

	for i, z := 1, 1.0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)

		z = math.Round(z*round) / round
		if math.Abs(guess-z) < 0.000001 {
			guess = z
			break
		}
		guess = z
	}

	return guess, nil
}

func main() {
	x := 2.0
	res, err := Sqrt(x)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Println("     Sqrt=", res)
	fmt.Println("math.Sqrt=", math.Sqrt(x))
}
