package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		result := add(100, 200)
		ch <- result
	}()
	fmt.Println("Result :", <-ch)
}

func add(x, y int) int {
	result := x + y
	return result
}
