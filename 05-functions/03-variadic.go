package main

import "fmt"

func main() {
	fmt.Println(sum(10, 20))
	fmt.Println(sum(10, 20, 30))
	fmt.Println(sum(10, 20, 30, 40))
	fmt.Println(sum(10))
	fmt.Println(sum())
}

func sum(list ...int) int {
	var result int
	for i := 0; i < len(list); i++ {
		result += list[i]
	}
	return result
}
