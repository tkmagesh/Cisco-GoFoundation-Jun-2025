/*
Make the execution of "checkPrime()" concurrent so that the number are not processed "sequentially"
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	var start, end int
	fmt.Println("Enter the range :")
	fmt.Scanln(&start, &end)
	printPrimes(start, end)
	fmt.Println("Done!")
}

func printPrimes(start, end int) {
	wg := &sync.WaitGroup{}
	for no := start; no <= end; no++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			go checkPrime(no)
		}()
	}
	wg.Wait()
}

func checkPrime(no int) {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return
		}
	}
	fmt.Println("Prime No :", no)
}
