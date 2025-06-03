package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("divisor cannot be zero")

func main() {
	// divisor := 7
	divisor := 0
	if q, r, err := divide(100, divisor); err == nil {
		fmt.Printf("quotient = %d and remainder = %d\n", q, r)
	} else if err == ErrDivideByZero {
		fmt.Println("Do not attempt to divide by zero")
	} else {
		fmt.Println("Error occurred :", err)
	}
}

// Named results

func divide(multiplier, divisor int) (quotient, remainder int, err error) { // quotient, remainder & err is automatically declared & initialized with the zero values
	if divisor == 0 {
		err = ErrDivideByZero
		return
	}
	quotient, remainder = multiplier/divisor, multiplier%divisor
	return
}
