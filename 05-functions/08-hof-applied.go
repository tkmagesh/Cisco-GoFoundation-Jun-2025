package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	add := profileWrapper(logWrapper(add))
	subtract := profileWrapper(logWrapper(subtract))

	add(100, 200)
	subtract(100, 200)

}

// ver 1.0
func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result :", x+y)
}

// ver 2.0
/*
// compose log capability
func logWrapper(oper func(int, int)) func(int, int) {
	return func(i1, i2 int) {
		log.Println("Operation started")
		oper(i1, i2)
		log.Println("Operation completed")
	}
}

// ver 3.0
func profileWrapper(oper func(int, int)) func(int, int) {
	return func(i1, i2 int) {
		start := time.Now()
		oper(i1, i2)
		elapsed := time.Since(start)
		log.Printf("Time taken : %v\n", elapsed)
	}
}
*/

// function signature as a 'type'
type OperationFn func(int, int)

// ver 2.0
// compose log capability
func logWrapper(oper OperationFn) OperationFn {
	return func(i1, i2 int) {
		log.Println("Operation started")
		oper(i1, i2)
		log.Println("Operation completed")
	}
}

// ver 3.0
func profileWrapper(oper OperationFn) OperationFn {
	return func(i1, i2 int) {
		start := time.Now()
		oper(i1, i2)
		elapsed := time.Since(start)
		log.Printf("Time taken : %v\n", elapsed)
	}
}
