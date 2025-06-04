/*
Make the execution of "checkPrime()" concurrent so that the number are not processed "sequentially"
*/

package main

import "fmt"

func main() {
	var start, end int
	fmt.Println("Enter the range :")
	fmt.Scanln(&start, &end)
	printPrimes(start, end)
}

func printPrimes(start, end int) {
	for no := start; no <= end; no++ {
		checkPrime(no)
	}
}

func checkPrime(no int) {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return
		}
	}
	fmt.Println("Prime No :", no)
}
