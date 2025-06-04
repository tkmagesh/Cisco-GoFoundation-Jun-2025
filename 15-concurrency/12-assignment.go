/*
Modify the below:
The generated prime number have to be printed in the main() function
IMPORTANT: the primes numbers have to be printed AS AND WHEN they are generated
*/

package main

import "fmt"

func main() {
	var start, end int
	fmt.Println("Enter the range :")
	fmt.Scanln(&start, &end)
	printPrimes(start, end)
}

func generatePrimes(start, end int) {
	for no := start; no <= end; no++ {
		checkPrime(no)
	}
}

func checkPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
