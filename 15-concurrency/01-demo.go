package main

import (
	"fmt"
	"time"
)

func main() {

	// scheduling the execution of f1() through the scheduler for f1() to be executed in future
	go f1()
	f2()

	// block the execution this function so that the scheduler can look for other goroutines that are scheduled and execute them (cooperative multitasking)

	// DO NOT do this (poor man's synchronization technique)
	time.Sleep(4 * time.Second)
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second) // simulate a time consuming operation
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
