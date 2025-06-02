package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hi!")
	}()

	func(userName string) {
		fmt.Printf("Hello, %s!\n", userName)
	}("Magesh")

	func(firstName, lastName string) {
		fmt.Printf("Hi %s %s, Have a good day!\n", firstName, lastName)
	}("Magesh", "Kuppan")

	fmt.Println("add(100,200) :", func(x, y int) int {
		return x + y
	}(100, 200))

	q, r := func(x, y int) (quotient, remainder int) { // quotient & remainder is automatically declared & initialized with the zero values
		quotient, remainder = x/y, x%y
		return
	}(100, 7)
	fmt.Printf("dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)

}
