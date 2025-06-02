package main

import "fmt"

func main() {
	sayHi()
	sayHello("Magesh")
	greet("Magesh", "Kuppan")
	fmt.Println("add(100,200) :", add(100, 200))

	/*
		fmt.Println("divide(100, 7) :")
		fmt.Println(divide(100, 7))
	*/

	// Using the "return results"

	q, r := divide(100, 7)
	fmt.Printf("dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)

	// Using ONLY ONE return result
	// Use "_" to receive a value that need to be ignored
	/*
		q, _ := divide(100, 7)
		fmt.Printf("dividing 100 by 7, quotient = %d\n", q)
	*/
}

func sayHi() {
	fmt.Println("Hi!")
}

func sayHello(userName string) {
	fmt.Printf("Hello, %s!\n", userName)
}

/*
func greet(firstName string, lastName string) {
	fmt.Printf("Hi %s %s, Have a good day!\n", firstName, lastName)
}
*/

func greet(firstName, lastName string) {
	fmt.Printf("Hi %s %s, Have a good day!\n", firstName, lastName)
}

func add(x, y int) int {
	return x + y
}

/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/

// Named results
func divide(x, y int) (quotient, remainder int) { // quotient & remainder is automatically declared & initialized with the zero values
	/*
		quotient = x / y
		remainder = x % y
	*/
	quotient, remainder = x/y, x%y
	return
}
