package calculator

import "fmt"

func init() {
	fmt.Println("[calculator] - init - (add.go)")
}

// public
func Add(x, y int) int {
	// opCount++
	opCount["Add"]++
	return x + y
}
