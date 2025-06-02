package main

import "fmt"

func main() {
	var n int
	n = 100

	// Value => Address (Pointer)
	// address of n
	fmt.Printf("&n = %p\n", &n)

	var nPtr *int
	nPtr = &n
	fmt.Printf("nPtr = %p\n", nPtr)

	// Address => Value (dereferencing)
	var x int
	x = *nPtr
	fmt.Println("x :", x)

	fmt.Println("n == *(&n) :", n == *(&n))
}
