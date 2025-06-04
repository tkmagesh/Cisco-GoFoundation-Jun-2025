/*
Modify the below:
The generated prime number have to be printed in the main() function
IMPORTANT: the primes numbers have to be printed AS AND WHEN they are generated
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	var start, end int
	ch := make(chan int)
	fmt.Println("Enter the range :")
	fmt.Scanln(&start, &end)
	go generatePrimes(start, end, ch)
	for primeNo := range ch {
		fmt.Println("Prime No :", primeNo)
	}
	fmt.Println("Done!")
}

func generatePrimes(start, end int, ch chan int) {
	wg := &sync.WaitGroup{}
	for no := start; no <= end; no++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if checkPrime(no) {
				ch <- no
			}
		}()
	}
	wg.Wait()
	close(ch)
}

func checkPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
