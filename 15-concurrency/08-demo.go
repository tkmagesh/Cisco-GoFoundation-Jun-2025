package main

import "fmt"

func main() {
	ch := make(chan int)
	go add(100, 200, ch)
	fmt.Println("Result :", <-ch)
}

func add(x, y int, ch chan int) {
	result := x + y
	ch <- result
}
