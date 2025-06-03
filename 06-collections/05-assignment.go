/*
Refactor the logic for generating prime nos to a "generatePrimes()" function
Write "IsPrime()" that checks if the given number is a prime number or not
Print the generated prime numbers in the main() function
*/

/*
Accept a range from the user and print all the prime numbers between the given range
*/

package main

import "fmt"

func main() {
LOOP:
	for no := 2; no <= 100; no++ {
		for i := 2; i <= (no / 2); i++ {
			if no%i == 0 {
				continue LOOP
			}
		}
		fmt.Println("Prime No :", no)
	}
}
