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
	var divisor int
	for {
		fmt.Println("Enter the divisor :")
		fmt.Scanln(&divisor)
		if q, r, err := divideFacade(100, divisor); err != nil {
			fmt.Printf("Error : %v.. try again!\n", err)
			continue
		} else {
			fmt.Printf("[main] quotient = %d and remainder = %d\n", q, r)
			break
		}
	}

}

// Use a facade to convert the panic into an error
func divideFacade(multiplier, divisor int) (quotient, remainder int, err error) {
	// handle the panic from the divide()
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	quotient, remainder = divide(multiplier, divisor)
	return
}

// 3rd party library
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
