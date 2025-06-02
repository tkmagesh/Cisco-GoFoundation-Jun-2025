/*
Refactor the below into small functions
Make sure each function has ONLY ONE reason to change (Single Responsibility Principle)
*/
package main

import "fmt"

func main() {
	var userChoice, n1, n2, result int
	for {

		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Println("5. Exit")
		fmt.Println("Enter your choice:")
		fmt.Scanln(&userChoice)
		if userChoice == 5 {
			break
		}
		if userChoice < 1 || userChoice > 5 {
			fmt.Println("Invalid Choice.. try again!")
			continue
		}
		fmt.Println("Enter the operands :")
		fmt.Scanln(&n1, &n2)
		switch userChoice {
		case 1:
			result = n1 + n2
		case 2:
			result = n1 - n2
		case 3:
			result = n1 * n2
		case 4:
			result = n1 / n2
		}
		fmt.Println("Result :", result)
	}
	fmt.Println("Thank you!")
}
