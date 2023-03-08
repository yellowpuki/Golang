package main

import (
	"errors"
	"fmt"
)

// DivisionError ...
type DivisionError struct {
	IntA int
	IntB int
	Msg  string
}

// Error ...
func (e *DivisionError) Error() string {
	return e.Msg
}

// doSomething ...
func doSomething() error {
	return errors.New("doSomething: something didn't work")
}

// Divide ...
func Divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, &DivisionError{
			Msg:  fmt.Sprintf("can't divide %d by zero", a),
			IntA: a, IntB: b,
		}
	}

	return a / b, nil
}

func main() {
	var v1, v2 int
	fmt.Scan(&v1, &v2)
	res, err := Divide(v1, v2)
	if err != nil {
		var divErr *DivisionError
		switch {
		case errors.As(err, &divErr):
			fmt.Printf("%d / %d is not mathematicly valid: %s\n",
				divErr.IntA, divErr.IntB, divErr.Error())
		default:
			fmt.Printf("unexpected divide error: %s\n", err)
		}
		return
	}

	fmt.Printf("%d / %d = %d\n", v1, v2, res)
}
