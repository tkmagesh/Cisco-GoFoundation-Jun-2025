package main

import "fmt"

func main() {
	var sayHi func()
	sayHi = func() {
		fmt.Println("Hi!")
	}
	if sayHi != nil {
		sayHi()
	}

	var sayHello func(string)
	sayHello = func(userName string) {
		fmt.Printf("Hello, %s!\n", userName)
	}
	sayHello("Magesh")

	var greet func(string, string)
	greet = func(firstName, lastName string) {
		fmt.Printf("Hi %s %s, Have a good day!\n", firstName, lastName)
	}
	greet("Magesh", "Kuppan")

	var add func(int, int) int
	add = func(x, y int) int {
		return x + y
	}
	fmt.Println("add(100,200) :", add(100, 200))

	var divide func(int, int) (int, int)
	divide = func(x, y int) (quotient, remainder int) { // quotient & remainder is automatically declared & initialized with the zero values
		quotient, remainder = x/y, x%y
		return
	}
	q, r := divide(100, 7)
	fmt.Printf("dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)

}
