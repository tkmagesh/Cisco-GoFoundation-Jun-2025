package main

import "fmt"

func main() {

	/*
		no := 22
		if no%2 == 0 {
			fmt.Printf("%d is an even number\n", no)
		} else {
			fmt.Printf("%d is an odd number\n", no)
		}
		fmt.Println(no)
	*/

	if no := 21; no%2 == 0 { // the 'no' scope is limited to 'if-else' block
		fmt.Printf("%d is an even number\n", no)
	} else {
		fmt.Printf("%d is an odd number\n", no)
	}
	// fmt.Println(no)

}
