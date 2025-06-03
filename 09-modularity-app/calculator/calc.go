package calculator

import "fmt"

// private
// var opCount int
var opCount map[string]int

// automatically executed whenever this package is imported and used
func init() {
	fmt.Println("[calculator] - init - (calc.go)")
	opCount = make(map[string]int)
}

/*
func OpCount() int {
	return opCount
}
*/

func OpCount() map[string]int {
	return opCount
}
