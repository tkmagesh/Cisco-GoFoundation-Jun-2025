package main

import (
	"fmt"

	"github.com/tkmagesh/cisco-gofoundation-jun-2025/09-modularity-app/calculator"
	// "github.com/tkmagesh/cisco-gofoundation-jun-2025/09-modularity-app/calculator/utils"

	// using an alias
	ut "github.com/tkmagesh/cisco-gofoundation-jun-2025/09-modularity-app/calculator/utils"
)

func run() {
	fmt.Println("application executed")
	fmt.Println(calculator.Add(100, 200))
	fmt.Println(calculator.Subtract(100, 200))
	fmt.Println("Operation count :", calculator.OpCount())
	// fmt.Println("is 21 even ? :", utils.IsEven(21))

	// using package alias
	fmt.Println("is 21 even ? :", ut.IsEven(21))
	fmt.Println("is 21 odd ? :", ut.IsOdd(21))
}
