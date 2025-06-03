package main

import (
	"errors"
	"fmt"
)

func main() {
	// divisor := 7
	divisor := 0
	if q, r, err := divide(100, divisor); err == nil {
		fmt.Printf("quotient = %d and remainder = %d\n", q, r)
	} else {
		fmt.Println("Error occurred :", err)
	}
}

/*
func divide(multiplier, divisor int) (int, int, error) {
	if divisor == 0 {
		e := errors.New("divisor cannot be zero")
		return 0, 0, e
	}
	quotient := multiplier / divisor
	remainder := multiplier % divisor
	return quotient, remainder, nil
}
*/

// Named results

func divide(multiplier, divisor int) (quotient, remainder int, err error) { // quotient, remainder & err is automatically declared & initialized with the zero values
	if divisor == 0 {
		err = errors.New("divisor cannot be zero")
		return
	}
	quotient, remainder = multiplier/divisor, multiplier%divisor
	return
}
