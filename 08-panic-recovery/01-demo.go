package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("divisor cannot be zero")

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("app panicked with error : ", e)
		}
		fmt.Println("Done!")
	}()
	divisor := 7

	// divisor := 0
	q, r := divide(100, divisor)
	fmt.Printf("[main] quotient = %d and remainder = %d\n", q, r)

}

func divide(multiplier, divisor int) (int, int) {

	fmt.Println("[divide] calculating quotient")
	if divisor == 0 {
		panic(ErrDivideByZero)
	}
	quotient := multiplier / divisor
	fmt.Println("[divide] calculating remainder")
	remainder := multiplier % divisor
	return quotient, remainder
}

// Named results
/*
func divide(multiplier, divisor int) (quotient, remainder int, err error) { // quotient, remainder & err is automatically declared & initialized with the zero values
	if divisor == 0 {
		err = errors.New("divisor cannot be zero")
		return
	}
	quotient, remainder = multiplier/divisor, multiplier%divisor
	return
} */
