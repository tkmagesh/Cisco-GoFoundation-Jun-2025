/*
Refactor the below into small functions
Make sure each function has ONLY ONE reason to change (Single Responsibility Principle)
*/
package main

import "fmt"

func main() {
	var userChoice, n1, n2 int
	for {
		getUserChoice(&userChoice)
		if userChoice == 5 {
			break
		}
		if userChoice < 1 || userChoice > 5 {
			fmt.Println("Invalid Choice.. try again!")
			continue
		}
		getOperands(&n1, &n2)
		processChoice(userChoice, n1, n2)
	}
	fmt.Println("Thank you!")
}

func processChoice(userChoice, n1, n2 int) {
	var result int
	switch userChoice {
	case 1:
		result = add(n1, n2)
	case 2:
		result = subtract(n1, n2)
	case 3:
		result = multply(n1, n2)
	case 4:
		result = divide(n1, n2)
	}
	fmt.Println("Result :", result)
}

func getUserChoice(userChoice *int) {
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("5. Exit")
	fmt.Println("Enter your choice:")
	fmt.Scanln(userChoice)
}

func getOperands(n1, n2 *int) {
	fmt.Println("Enter the operands :")
	fmt.Scanln(&n1, &n2)
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return x / y
}
