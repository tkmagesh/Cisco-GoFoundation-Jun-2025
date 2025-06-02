package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for range 20 {
		fn := getFn()
		fn()
		time.Sleep(500 * time.Millisecond)
	}
}

func getFn() func() {
	switch n := rand.Intn(20); {
	case n%2 == 0:
		return f1
	case n%3 == 0:
		return f2
	default:
		return func() {
			fmt.Println("anonymous fn invoked")
		}
	}
}
func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
