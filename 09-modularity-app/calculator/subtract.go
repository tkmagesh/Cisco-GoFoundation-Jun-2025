package calculator

import "fmt"

func init() {
	fmt.Println("[calculator] - init - (subtract.go)")
}

func Subtract(x, y int) int {
	// opCount++
	opCount["Subtract"]++
	return x - y
}
